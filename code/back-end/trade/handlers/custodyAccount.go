package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lightningnetwork/lnd/lnrpc"
	"net/http"
	"time"
	"trade/middleware"
	"trade/models"
	"trade/services"
	"trade/services/custodyAccount"
)

const (
	BILL_TYPE_RECHARGE = 0
	BILL_TYPE_PAYMENT  = 1
)

const (
	AWAY_IN  = 0
	AWAY_OUT = 1
)

const (
	UNIT_SATOSHIS = 0
)

const (
	PAY_UNKNOWN = 0
	PAY_SUCCESS = 1
	PAY_FAILED  = 2
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
	//TODO: 判断申请金额是否超过通道余额
	memo := "testmemo"
	amount := int64(1000)
	//生成一张发票
	invoiceRequest, err := custodyAccount.ApplyInvoice(accountCode, amount, memo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//获取发票信息
	info, _ := custodyAccount.FindInvoice(invoiceRequest.RHash)

	//构建invoice对象
	var invoiceModel models.Invoice
	invoiceModel.UserID = userId
	invoiceModel.Invoice = invoiceRequest.PaymentRequest
	invoiceModel.AccountID = &accountId
	invoiceModel.Amount = float64(amount)

	invoiceModel.Status = int16(info.State)
	template := time.Unix(info.CreationDate, 0)
	invoiceModel.CreateDate = &template
	expiry := int(info.Expiry)
	invoiceModel.Expiry = &expiry

	//写入数据库
	err = middleware.DB.Create(&invoiceModel).Error
	//回信息，规范状态码
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoiceModel": invoiceModel})
}

// PayInvoice CustodyAccountzhi付款发票
func PayInvoice(c *gin.Context) {

	//TODO: 获取登录用户信息
	//userId := uint(0)
	accountId := uint(0)
	accountCode := "8bc6754444ab8020"

	//TODO: 校验发票信息
	invoice := "lntb30u1pnysgl9pp5tphfml5fayvw67vpd72y3aavekuzs6jv04ga7yvw7hhsy2l96cvsdqqcqzzsxqyz5vqsp55nlk7zltyecpktehpdgmknjaelyztwrnpqy6cy99447g83hhghhs9qyyssqm4ldvk854ap005x2fq72tp635xz5gzkuw9puz8q02pcs9qk30938rer3ah9avhpvpldv9mhwuqhuqa4gvd4ezeqqx955q9pgfvrlarqq8nnt3f"
	feeLimit := int64(1000)
	//检查数据库中是否有该发票的记录
	a, err := services.GenericQueryByObject(&models.Balance{
		Invoice: &invoice,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if len(a) > 0 {
		for _, v := range a {
			if v.Status == PAY_SUCCESS {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票已被支付，请勿重复支付"})
				return
			}
			if v.Status == PAY_UNKNOWN {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票锁定中，请稍后再试"})
				return
			}
		}
	}

	//TODO: 判断账户余额是否足够
	info, err := custodyAccount.DecodeInvoice(invoice)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	//TODO: 判断发票支付条件
	userbalance := int64(100000)

	if info.NumSatoshis > userbalance {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "余额不足"})
		return
	}

	// 支付发票
	payment, err := custodyAccount.PayInvoice(accountCode, invoice, feeLimit)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// 写入数据库
	var balanceModel models.Balance
	balanceModel.AccountId = accountId
	balanceModel.BillType = BILL_TYPE_PAYMENT
	balanceModel.Away = AWAY_OUT
	balanceModel.Amount = float64(payment.ValueSat)
	balanceModel.Unit = UNIT_SATOSHIS
	balanceModel.Invoice = &invoice
	balanceModel.PaymentHash = &payment.PaymentHash
	if payment.Status == lnrpc.Payment_SUCCEEDED {
		balanceModel.Status = PAY_SUCCESS
	} else if payment.Status == lnrpc.Payment_FAILED {
		balanceModel.Status = PAY_FAILED
	} else {
		balanceModel.Status = PAY_UNKNOWN
	}
	err = middleware.DB.Create(&balanceModel).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}
