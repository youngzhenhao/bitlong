package api

import "log"

type WalletStatus string

const (
	WALLET_EMPTY    WalletStatus = "[无钱包]"
	WALLET_UNLOCKED WalletStatus = "[已解锁]"
	WALLET_LOCKED   WalletStatus = "[未解锁]"
)

// LndWorkflowCreateWallet 测试创建钱包
func LndWorkflowCreateWallet(password string) bool {
	log.Println(WALLET_EMPTY, "开始生成随机助记词用以创建钱包...")
	seed := GenSeed()
	log.Println("助记词已生成:\n", seed)
	log.Println("使用该助记词和密码初始化钱包中...")
	result := InitWallet(seed, password)
	if result {
		log.Println(WALLET_UNLOCKED, "钱包创建成功！")
	}
	return result
}

func LndWorkflowUnlockWallet(password string) bool {
	log.Println(WALLET_LOCKED, "正在解锁...")
	result := UnlockWallet(password)
	if result {
		log.Println(WALLET_UNLOCKED, "钱包解锁成功！")
	}
	return result
}

func LndWorkflowChangeWalletPassword() {

}
