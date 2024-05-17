package custodyAccount

import (
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"os"
	"path/filepath"
	"trade/config"
)

// CreateCustodyAccount 创建托管账户并保持马卡龙文件
func CreateCustodyAccount(label string) (*litrpc.Account, error) {
	//根据用户信息创建托管账户
	account, macaroon, err := accountCreate(0, 0, label)
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

	//返回托管账户信息
	return account, nil
}

// Update TODO: 托管账户充值
func Update(id string, amount int64) error {
	//更变托管账户余额
	_, err := accountUpdate(id, amount, -1)
	if err != nil {
		return err
	}
	//TODUO: 返回充值结果
	return err
}

// QueryCustodyAccount TODO: 托管账户查询
func QueryCustodyAccount() error {
	//TODO: 获取托管账户ID
	id := "test"
	//从节点查询账户信息
	account, err := accountInfo(id)
	if err != nil {
		return err
	}
	//TODO: 查询数据库相关信息

	//TODO: 组合并返回账户信息
	fmt.Println(account)
	return err
}

// DeleteCustodianAccount TODO: 托管账户删除
func DeleteCustodianAccount() error {
	//TODO: 获取托管账户ID
	id := "test"
	//删除Lit节点托管账户
	err := accountRemove(id)
	//TODO: 更新数据库相关信息

	//TODO: 返回删除结果

	return err
}

// ApplyInvoice TODO: 使用指定账户申请一张发票
func ApplyInvoice(id string, amount int64, memo string) (string, error) {
	//获取马卡龙路径
	var macaroonFile string
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir

	if id == "admin" {
		macaroonFile = config.GetConfig().ApiConfig.Lnd.MacaroonPath
	} else {
		macaroonFile = filepath.Join(macaroonDir, id+".macaroon")
	}
	//调用Lit节点发票申请接口
	return invoiceCreate(amount, memo, macaroonFile)
}

// PayInvoice TODO: 使用指定账户支付发票
func PayInvoice() {
	//TODO: 验证账户信息，发票信息，发票金额是否正确
	//TODO: 获取托管账户macaroon
	//TODO: 调用Lit节点发票支付接口
	//TODO: 更新数据库相关信息
	//TODO: 返回支付结果
}

// DecodeInvoice  解析发票信息
func DecodeInvoice(invoice string) (*lnrpc.PayReq, error) {
	return invoiceDecode(invoice)
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
