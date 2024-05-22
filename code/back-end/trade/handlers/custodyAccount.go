package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	_, err = services.QueryCustodyAccount(account.UserAccountCode)
	amount := uint64(1)

	//TODO: 更新托管账户余额
	err = services.UpdateCustodyAccount(account, services.AWAY_IN, amount)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//TODO: 写入数据库

	//TODO: 返回信息
}

// ApplyInvoiceCA CustodyAccount开具发票
func ApplyInvoiceCA(c *gin.Context) {
	// 获取登录用户信息
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
	apply := services.ApplyRequest{}
	if err = c.ShouldBindJSON(&apply); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if apply.Amount <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发票信息不合法"})
		return
	}
	//生成一张发票
	invoiceRequest, err := services.ApplyInvoice(user, account, &apply)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoiceModel": invoiceRequest})
}

// PayInvoice CustodyAccount付款发票
func PayInvoice(c *gin.Context) {
	// 获取登录用户信息
	userName := c.MustGet("username").(string)
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	// 选择托管账户
	account, err := services.ReadAccountByUserId(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if account.UserAccountCode == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未找到账户信息"})
		return
	}

	//获取支付发票请求
	pay := services.PayInvoiceRequest{}
	if err := c.ShouldBindJSON(&pay); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	// 支付发票
	payment, err := services.PayInvoice(account, &pay)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}
