package services

import (
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"os"
	"path/filepath"
	"sync"
	"trade/config"
	"trade/models"
	"trade/services/rpc"
)

var mutex sync.Mutex

// CreateCustodyAccount 创建托管账户并保持马卡龙文件
func CreateCustodyAccount(user *models.User) (*litrpc.Account, error) {
	//根据用户信息创建托管账户
	account, macaroon, err := rpc.AccountCreate(0, 0)
	if err != nil {
		return nil, err
	}

	//构建马卡龙存储路径
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir
	if _, err = os.Stat(macaroonDir); os.IsNotExist(err) {
		err = os.MkdirAll(macaroonDir, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目标文件夹 %s 失败: %v\n", macaroonDir, err)
			return nil, err
		}
	}
	macaroonFile := filepath.Join(macaroonDir, account.Id+".macaroon")

	//存储马卡龙信息
	err = saveMacaroon(macaroon, macaroonFile)
	if err != nil {
		return nil, err
	}

	//构建账户对象
	var accountModel models.Account
	accountModel.UserName = user.Username
	accountModel.UserId = user.ID
	accountModel.UserAccountCode = account.Id
	accountModel.Label = &account.Label
	accountModel.Status = 1

	//写入数据库
	mutex.Lock()
	defer mutex.Unlock()
	err = CreateAccount(&accountModel)

	//返回托管账户信息
	return account, nil
}

// Update  托管账户充值
func Update(id string, amount int64) error {
	//更变托管账户余额
	_, err := rpc.AccountUpdate(id, amount, -1)
	if err != nil {
		return err
	}
	//TODUO: 返回充值结果
	return err
}

// QueryCustodyAccount  托管账户查询
func QueryCustodyAccount(accountCode string) (*litrpc.Account, error) {
	return rpc.AccountInfo(accountCode)
}

// DeleteCustodianAccount 托管账户删除
func DeleteCustodianAccount() error {
	//TODO: 获取托管账户ID
	id := "test"
	//删除Lit节点托管账户
	err := rpc.AccountRemove(id)
	//TODO: 更新数据库相关信息

	//TODO: 返回删除结果

	return err
}

// ApplyInvoice 使用指定账户申请一张发票
func ApplyInvoice(accountCode string, amount int64, memo string) (*lnrpc.AddInvoiceResponse, error) {
	//获取马卡龙路径
	var macaroonFile string
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir

	if accountCode == "admin" {
		macaroonFile = config.GetConfig().ApiConfig.Lnd.MacaroonPath
	} else {
		macaroonFile = filepath.Join(macaroonDir, accountCode+".macaroon")
	}
	if macaroonFile == "" {
		return nil, fmt.Errorf("macaroon file not found")
	}
	//调用Lit节点发票申请接口
	return rpc.InvoiceCreate(amount, memo, macaroonFile)
}

// PayInvoice 使用指定账户支付发票
func PayInvoice(accountCode string, invoice string, feeLimit int64) (*lnrpc.Payment, error) {
	//获取马卡龙路径
	var macaroonFile string
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir

	if accountCode == "admin" {
		macaroonFile = config.GetConfig().ApiConfig.Lnd.MacaroonPath
	} else {
		macaroonFile = filepath.Join(macaroonDir, accountCode+".macaroon")
	}
	if macaroonFile == "" {
		return nil, fmt.Errorf("macaroon file not found")
	}

	return rpc.InvoicePay(macaroonFile, invoice, feeLimit)
}

// DecodeInvoice  解析发票信息
func DecodeInvoice(invoice string) (*lnrpc.PayReq, error) {
	return rpc.InvoiceDecode(invoice)
}

// FindInvoice 查询节点内部发票
func FindInvoice(rHash []byte) (*lnrpc.Invoice, error) {
	return rpc.InvoiceFind(rHash)
}

// TrackPayment 跟踪支付状态
func TrackPayment(paymentHash string) (*lnrpc.Payment, error) {
	return rpc.PaymentTrack(paymentHash)
}

// saveMacaroon 保存macaroon字节切片到指定文件
func saveMacaroon(macaroon []byte, macaroonFile string) error {
	file, err := os.OpenFile(macaroonFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	// 将字节切片写入指定位置
	data := macaroon
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
