package custodyAccount

import (
	"fmt"
	"os"
)

// TODO: 创建托管账户
func CreateCustodyAccount() {
	// TODO: 读取请求的用户信息
	label := "test user 1"
	//根据用户信息创建托管账户
	_, macaroon, err := accountCreate(0, 0, label)
	if err != nil {
		fmt.Println(err)
	}
	//TODO: 创建存储托管账户信息对象

	//TODO: 存储马卡龙信息
	fmt.Print(macaroon)
	//TODO: 存储托管账户信息
	//TODO: 返回托管账户信息
}

// TODO: 托管账户充值
func Recharge() error {
	//TODO:获取托管账户信息和充值额度,计算充值后账户余额
	id := "test"
	amount := int64(1000000000)
	//TODO: 验证交易

	//更变托管账户余额
	_, err := accountUpdate(id, amount, -1)
	if err != nil {
		return err
	}
	//TODO: 更新数据库

	//TODUO: 返回充值结果
	return err
}

// TODO: 托管账户查询
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

// TODO: 托管账户删除
func DeleteCustodianAccount() error {
	//TODO: 获取托管账户ID
	id := "test"
	//删除Lit节点托管账户
	err := accountRemove(id)
	//TODO: 更新数据库相关信息

	//TODO: 返回删除结果

	return err
}

// TODO: 使用指定账户申请一张发票
func ApplyInvoice() {
	//TODO: 验证账户信息，发票信息，发票金额是否正确
	//TODO: 获取托管账户macaroon
	//TODO: 调用Lit节点发票申请接口
	//TODO: 存储发票信息
	//TODO: 返回发票信息
}

// TODO: 使用指定账户支付发票
func PayInvoice() {
	//TODO: 验证账户信息，发票信息，发票金额是否正确
	//TODO: 获取托管账户macaroon
	//TODO: 调用Lit节点发票支付接口
	//TODO: 更新数据库相关信息
	//TODO: 返回支付结果
}

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
