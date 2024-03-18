package tx

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

type Unspent struct {
	TxHash string
	Index  uint32
	Value  int64
}
type receiver struct {
	amount  int64
	address string
}

func CreateTxs(prikey string, toAddress string, amount, fee int64, unspent []*Unspent, isWitness bool) (string, error) {
	var totalBal int64
	for _, v := range unspent {
		totalBal += v.Value
	}
	wif, err := btcutil.DecodeWIF(prikey)
	if err != nil {
		return "", err
	}
	p2pkhAddr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	recivers := []*receiver{
		{amount: amount, address: toAddress},
	}
	backBal := totalBal - amount - fee
	if backBal > 0 { // 找零
		recivers = append(recivers, &receiver{
			amount:  backBal,
			address: p2pkhAddr.EncodeAddress(),
		})
	}

	tx := wire.NewMsgTx(wire.TxVersion)

	// 构造输入
	for _, v := range unspent {
		hash, err := chainhash.NewHashFromStr(v.TxHash)
		if err != nil {
			return "", err
		}
		outPoint := wire.NewOutPoint(hash, v.Index)
		tx.AddTxIn(wire.NewTxIn(outPoint, nil, nil))
	}
	// 构造输出
	for _, v := range recivers {
		addr, err := btcutil.DecodeAddress(v.address, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		script, err := txscript.PayToAddrScript(addr)
		if err != nil {
			return "", err
		}
		tx.AddTxOut(wire.NewTxOut(v.amount, script))
	}
	// 判断隔离见证签名
	if isWitness {
		for i, _ := range tx.TxIn {
			pubKeyHash := btcutil.Hash160(wif.PrivKey.PubKey().SerializeCompressed())
			p2wkhAddr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
			if err != nil {
				fmt.Println(err)
			}
			//生成见证脚本Script
			witnessProgram, err := txscript.PayToAddrScript(p2wkhAddr)
			if err != nil {
				fmt.Println(err)
			}
			//交易哈希
			hashCache := txscript.NewTxSigHashes(tx, nil)
			//见证脚本交易签名 返回txinwitness
			witnessScript, err := txscript.WitnessSignature(tx, hashCache, 0, totalBal, witnessProgram, txscript.SigHashAll, wif.PrivKey, true)
			tx.TxIn[i].Witness = witnessScript
			//填充隔离见证scriptSig
			bldr := txscript.NewScriptBuilder().AddData(witnessProgram)
			SignatureScript, err := bldr.Script()
			tx.TxIn[i].SignatureScript = SignatureScript
		}
	} else {
		for i, _ := range tx.TxIn {
			addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &chaincfg.MainNetParams)
			if err != nil {
				return "", err
			}
			script, err := txscript.PayToAddrScript(addr)
			if err != nil {
				return "", err
			}
			tx.TxIn[i].SignatureScript, err = txscript.SignTxOutput(&chaincfg.MainNetParams, tx, i, script, txscript.SigHashAll,
				txscript.KeyClosure(func(addr btcutil.Address) (*btcec.PrivateKey,
					bool, error) {
					return wif.PrivKey, true, nil
				}), nil, nil)
			if err != nil {
				return "", err
			}
		}
	}
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err = tx.Serialize(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

// g版
func CreateTx(prikey string, toAddress string, amount, fee int64, unspent []*Unspent, isWitness bool) (string, error) {
	var totalBal int64
	for _, v := range unspent {
		totalBal += v.Value
	}
	wif, err := btcutil.DecodeWIF(prikey)
	if err != nil {
		return "", err
	}
	p2pkhAddr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	recivers := []*receiver{
		{amount: amount, address: toAddress},
	}
	backBal := totalBal - amount - fee
	if backBal > 0 { // 找零
		recivers = append(recivers, &receiver{
			amount:  backBal,
			address: p2pkhAddr.EncodeAddress(),
		})
	}
	// 创建事务
	tx := wire.NewMsgTx(wire.TxVersion)
	// 添加输入
	for _, u := range unspent {
		hash, err := chainhash.NewHashFromStr(u.TxHash)
		if err != nil {
			return "", err
		}
		outPoint := wire.NewOutPoint(hash, u.Index)
		txIn := wire.NewTxIn(outPoint, nil, nil)
		tx.AddTxIn(txIn)
	}
	// 添加输出（包括找零）
	for _, v := range recivers {
		addr, err := btcutil.DecodeAddress(v.address, &chaincfg.MainNetParams)
		if err != nil {
			return "", err
		}
		script, err := txscript.PayToAddrScript(addr)
		if err != nil {
			return "", err
		}
		tx.AddTxOut(wire.NewTxOut(v.amount, script))
	}
	// 对每个输入进行签名
	for i, u := range unspent {
		var script []byte
		if isWitness {
			addr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &chaincfg.MainNetParams)
			if err != nil {
				return "", err
			}
			script, err = txscript.PayToAddrScript(addr)
			if err != nil {
				return "", err
			}
		} else {
			addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &chaincfg.MainNetParams)
			if err != nil {
				return "", err
			}
			script, err = txscript.PayToAddrScript(addr)
			if err != nil {
				return "", err
			}
		}
		if isWitness {
			// 对于隔离见证交易的签名
			sigHashes := txscript.NewTxSigHashes(tx, nil)
			signature, err := txscript.WitnessSignature(tx, sigHashes, i, u.Value, script, txscript.SigHashAll, wif.PrivKey, true)
			if err != nil {
				return "", err
			}
			tx.TxIn[i].Witness = signature
		} else {
			// 对于非隔离见证交易的签名
			signatureScript, err := txscript.SignatureScript(tx, i, script, txscript.SigHashAll, wif.PrivKey, true)
			if err != nil {
				return "", err
			}
			tx.TxIn[i].SignatureScript = signatureScript
		}
	}
	// 序列化事务
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return "", err
	}

	return hex.EncodeToString(buf.Bytes()), nil
}
