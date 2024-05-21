package handlers

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lightningnetwork/lnd/lnrpc"
	"net/http"
	"sync"
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

var mutex sync.Mutex

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
	cstAccount, err := custodyAccount.CreateCustodyAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//构建账户对象
	var accountModel models.Account
	accountModel.UserName = user.Username
	accountModel.UserId = user.ID
	accountModel.UserAccountCode = cstAccount.Id
	accountModel.Label = &cstAccount.Label
	accountModel.Status = 1
	//写入数据库
	mutex.Lock()
	defer mutex.Unlock()
	err = services.CreateAccount(&accountModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accountModel": accountModel})
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
	acc, err := custodyAccount.QueryCustodyAccount(account.UserAccountCode)
	amount := int64(1)

	//TODO: 更新托管账户余额
	err = custodyAccount.Update(acc.Id, amount)
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

	//TODO: 判断申请金额是否超过通道余额
	apply := struct {
		Amount int64  `json:"amount"`
		Memo   string `json:"memo"`
	}{}
	if err := c.ShouldBindJSON(&apply); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	//生成一张发票
	invoiceRequest, err := custodyAccount.ApplyInvoice(account.UserAccountCode, apply.Amount, apply.Memo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	//获取发票信息
	info, _ := custodyAccount.FindInvoice(invoiceRequest.RHash)

	//构建invoice对象
	var invoiceModel models.Invoice
	invoiceModel.UserID = user.ID
	invoiceModel.Invoice = invoiceRequest.PaymentRequest
	invoiceModel.AccountID = &account.ID
	invoiceModel.Amount = float64(info.Value)

	invoiceModel.Status = int16(info.State)
	template := time.Unix(info.CreationDate, 0)
	invoiceModel.CreateDate = &template
	expiry := int(info.Expiry)
	invoiceModel.Expiry = &expiry

	//写入数据库
	mutex.Lock()
	defer mutex.Unlock()
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
	userName := c.MustGet("username").(string)
	user, err := services.ReadUserByUsername(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}
	account, _ := services.ReadAccountByUserId(user.ID)

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
			if v.State == PAY_SUCCESS {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票已被支付，请勿重复支付"})
				return
			}
			if v.State == PAY_UNKNOWN {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "发票锁定中，请稍后再试"})
				return
			}
		}
	}

	// 判断账户余额是否足够
	info, err := custodyAccount.DecodeInvoice(pay.Invoice)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	userBalance, err := custodyAccount.QueryCustodyAccount(account.UserAccountCode)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if info.NumSatoshis > userBalance.CurrentBalance {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "余额不足"})
		return
	}

	// 支付发票
	payment, err := custodyAccount.PayInvoice(account.UserAccountCode, pay.Invoice, pay.FeeLimit)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// 写入数据库
	var balanceModel models.Balance
	balanceModel.AccountId = account.ID
	balanceModel.BillType = BILL_TYPE_PAYMENT
	balanceModel.Away = AWAY_OUT
	balanceModel.Amount = float64(payment.ValueSat)
	balanceModel.Unit = UNIT_SATOSHIS
	balanceModel.Invoice = &pay.Invoice
	balanceModel.PaymentHash = &payment.PaymentHash
	if payment.Status == lnrpc.Payment_SUCCEEDED {
		balanceModel.State = PAY_SUCCESS
	} else if payment.Status == lnrpc.Payment_FAILED {
		balanceModel.State = PAY_FAILED
	} else {
		balanceModel.State = PAY_UNKNOWN
	}
	mutex.Lock()
	defer mutex.Unlock()
	err = middleware.DB.Create(&balanceModel).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

// PollPayment 遍历所有未确认的发票，轮询支付状态
func PollPayment() {
	//查询数据库，获取所有未确认的支付
	params := services.QueryParams{
		"State": PAY_UNKNOWN,
	}
	a, err := services.GenericQuery(&models.Balance{}, params)
	if err != nil {
		fmt.Println(err)
	}
	if len(a) > 0 {
		for _, v := range a {
			temp, err := custodyAccount.TrackPayment(*v.PaymentHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if temp.Status == lnrpc.Payment_SUCCEEDED {
				v.State = PAY_SUCCESS
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					fmt.Println(err)
				}
			} else if temp.Status == lnrpc.Payment_FAILED {
				v.State = PAY_FAILED
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// PollInvoice 遍历所有未支付的发票，轮询发票状态
func PollInvoice() {
	//查询数据库，获取所有未支付的发票
	params := services.QueryParams{
		"Status": lnrpc.Invoice_OPEN,
	}
	a, err := services.GenericQuery(&models.Invoice{}, params)
	if err != nil {
		fmt.Println(err)
	}
	if len(a) > 0 {
		for _, v := range a {
			invoice, err := custodyAccount.DecodeInvoice(v.Invoice)
			if err != nil {
				fmt.Println(err)
				continue
			}
			rHash, err := hex.DecodeString(invoice.PaymentHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp, err := custodyAccount.FindInvoice(rHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if int16(temp.State) != v.Status {
				v.Status = int16(temp.State)
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

type CronHandler struct {
}

func (sm *CronHandler) PollPaymentCron() {
	fmt.Println("start cron job: PollPayment")
	PollPayment()
}
func (sm *CronHandler) PollInvoiceCron() {
	fmt.Println("start cron job: PollInvoice")
	PollInvoice()
}
