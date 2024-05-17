package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"trade/middleware"
	"trade/models"
	"trade/services"
	"trade/services/custodyAccount"
)

// CreateCustodyAccount 创建托管账户
func CreateCustodyAccount(c *gin.Context) {
	//var creds models.Account
	//if err := c.ShouldBindJSON(&creds); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//TODO: 获取登录用户信息
	//TODO: 校验登录用户信息
	userName := "testname"
	userId := uint(0)
	Lebel := "testlabel"

	//创建账户
	cstAccount, err := custodyAccount.CreateCustodyAccount(Lebel)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//构建账户对象
	var accountModel models.Account
	accountModel.UserName = userName
	accountModel.UserId = userId
	accountModel.UserAccountCode = cstAccount.Id
	accountModel.Label = &cstAccount.Label
	accountModel.Status = 1
	//写入数据库
	err = services.CreateAccount(&accountModel)
	//TODO: 返回信息，规范状态码
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accountModel": accountModel})
}

// UpdateCustodyAccount 更新托管账户余额
func UpdateCustodyAccount(c *gin.Context) {
	//TODO: 获取登录用户信息
	//TODO: 获取账户余额更新信息
	id := "testid"
	amount := int64(12345)

	//TODO: 更新托管账户余额
	err := custodyAccount.Update(id, amount)
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
	userId := uint(0)
	accountId := uint(0)
	accountCode := "8bc6754444ab8020"
	//TODO: 判断申请金额是否超过节点余额
	memo := "testmemo"
	amount := int64(1000)
	//生成一张发票
	invoice, err := custodyAccount.ApplyInvoice(accountCode, amount, memo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//构建invoice对象
	var invoiceModel models.Invoice
	invoiceModel.UserID = userId
	invoiceModel.Invoice = invoice
	invoiceModel.AccountID = &accountId
	invoiceModel.Amount = float64(amount)
	invoiceModel.Status = 1

	//获取发票信息
	info, _ := custodyAccount.DecodeInvoice(invoice)
	if info != nil {
		template := time.Unix(info.Timestamp, 0)
		invoiceModel.CreateDate = &template
		expiry := int(info.Expiry)
		invoiceModel.Expiry = &expiry
	}
	//写入数据库
	err = middleware.DB.Create(&invoiceModel).Error
	//回信息，规范状态码
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"invoiceModel": invoiceModel})
}

// PayInvoice CustodyAccountzhi付款发票
func PayInvoice(c *gin.Context) {
	//TODO: 获取登录用户信息

	//TODO: 校验发票信息

	//TODO: 判断发票支付条件

	//TODO: 支付发票

	//TODO: 写入数据库

	//TODO: 返回信息，规范状态码
}
