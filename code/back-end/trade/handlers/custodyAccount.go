package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"trade/models"
	"trade/services"
)

// CreateCustodyAccount 创建托管账户
func CreateCustodyAccount(c *gin.Context) {
	// 获取登录用户信息
	userName := c.MustGet("username").(string)

	// 校验登录用户信息
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	// 判断用户是否已经创建账户
	accounts, _ := services.ReadAccountByUserIds(user.ID)
	if len(accounts) > 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "托管账户已存在"})
		return
	}

	//创建账户
	cstAccount, err := services.CreateCustodyAccount(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountModel": cstAccount})
}

// UpdateCustodyAccount 更新托管账户余额
func UpdateCustodyAccount(c *gin.Context) {

	//TODO: 获取登录用户信息
	userName := c.MustGet("username").(string)
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	account, _ := services.ReadAccountByUserId(user.ID)

	//TODO: 获取账户余额更新信息
	acc, err := services.QueryCustodyAccount(account.UserAccountCode)
	amount := int64(1)

	//TODO: 更新托管账户余额
	err = services.Update(acc.Id, amount)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//TODO: 写入数据库

	//TODO: 返回信息
}

// ApplyInvoiceCA CustodyAccount开具发票
func ApplyInvoiceCA(c *gin.Context) {
	//TODO: 获取登录用户信息
	userName := c.MustGet("username").(string)
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	account, err := services.ReadAccountByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if account.UserAccountCode == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未找到账户信息"})
		//TODO 为用户创建托管账户
		return
	}

	//TODO: 判断申请金额是否超过通道余额,检查申请内容是否合法
	apply := struct {
		Amount int64  `json:"amount"`
		Memo   string `json:"memo"`
	}{}
	if err := c.ShouldBindJSON(&apply); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if apply.Amount <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发票信息不合法"})
		return
	}
	//生成一张发票
	invoiceRequest, err := services.ApplyInvoice(user, account, apply.Amount, apply.Memo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoiceModel": invoiceRequest})
}

// PayInvoice CustodyAccount付款发票
func PayInvoice(c *gin.Context) {

	//TODO: 获取登录用户信息
	userName := c.MustGet("username").(string)
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	account, err := services.ReadAccountByUserId(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if account.UserAccountCode == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未找到账户信息"})
		return
	}

	//TODO: 校验发票信息
	pay := struct {
		Invoice  string `json:"invoice"`
		FeeLimit int64  `json:"feeLimit"`
	}{}
	if err := c.ShouldBindJSON(&pay); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	//检查数据库中是否有该发票的记录
	a, err := services.GenericQueryByObject(&models.Balance{
		Invoice: &pay.Invoice,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if len(a) > 0 {
		for _, v := range a {
			if v.State == services.PAY_SUCCESS {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票已被支付，请勿重复支付"})
				return
			}
			if v.State == services.PAY_UNKNOWN {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票锁定中，请稍后再试"})
				return
			}
		}
	}

	// 判断账户余额是否足够
	info, err := services.DecodeInvoice(pay.Invoice)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	userBalance, err := services.QueryCustodyAccount(account.UserAccountCode)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if info.NumSatoshis > userBalance.CurrentBalance {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "余额不足"})
		return
	}

	// 支付发票
	payment, err := services.PayInvoice(account, pay.Invoice, pay.FeeLimit)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// 写入数据库

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}
