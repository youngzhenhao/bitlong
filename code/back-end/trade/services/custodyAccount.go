package services

import (
	"encoding/hex"
	"fmt"
	"github.com/lightninglabs/lightning-terminal/litrpc"
	"github.com/lightningnetwork/lnd/lnrpc"
	"os"
	"path/filepath"
	"sync"
	"time"
	"trade/config"
	"trade/middleware"
	"trade/models"
	"trade/services/servicesrpc"
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

// CreateCustodyAccount 创建托管账户并保持马卡龙文件
func CreateCustodyAccount(user *models.User) (*models.Account, error) {
	//根据用户信息创建托管账户
	account, macaroon, err := servicesrpc.AccountCreate(0, 0)
	if err != nil {
		CUST.Error(err.Error())
		return nil, err
	}

	//构建马卡龙存储路径
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir
	if _, err = os.Stat(macaroonDir); os.IsNotExist(err) {
		err = os.MkdirAll(macaroonDir, os.ModePerm)
		if err != nil {
			CUST.Error(fmt.Sprintf("创建目标文件夹 %s 失败: %v\n", macaroonDir, err))
			return nil, err
		}
	}
	macaroonFile := filepath.Join(macaroonDir, account.Id+".macaroon")

	//存储马卡龙信息
	err = saveMacaroon(macaroon, macaroonFile)
	if err != nil {
		CUST.Error(err.Error())
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
	if err != nil {
		CUST.Error(err.Error())
		return nil, err
	}
	//返回托管账户信息
	return &accountModel, nil
}

// Update  托管账户更新
func UpdateCustodyAccount(account *models.Account, away uint, balance uint64) error {
	acc, err := servicesrpc.AccountInfo(account.UserAccountCode)
	if err != nil {
		return err
	}
	var amount int64

	switch away {
	case AWAY_IN:
		amount = acc.CurrentBalance + int64(balance)
	case AWAY_OUT:
		amount = acc.CurrentBalance - int64(balance)
	default:
		return fmt.Errorf("away error")
	}

	//更变托管账户余额
	_, err = servicesrpc.AccountUpdate(account.UserAccountCode, amount, -1)
	if err != nil {
		return err
	}
	//构建数据库存储对象
	ba := models.Balance{}
	ba.AccountId = account.ID
	ba.Amount = float64(balance)
	ba.Unit = UNIT_SATOSHIS
	ba.BillType = BILL_TYPE_PAYMENT
	ba.Away = away
	ba.State = PAY_SUCCESS
	ba.Invoice = nil
	ba.PaymentHash = nil
	//更新数据库
	mutex.Lock()
	defer mutex.Unlock()
	err = middleware.DB.Create(&ba).Error
	if err != nil {
		CUST.Error(err.Error())
		return err
	}
	return nil
}

// QueryCustodyAccount  托管账户查询
func QueryCustodyAccount(accountCode string) (*litrpc.Account, error) {
	return servicesrpc.AccountInfo(accountCode)
}

// DeleteCustodianAccount 托管账户删除
func DeleteCustodianAccount() error {
	//TODO: 获取托管账户ID
	id := "test"
	//删除Lit节点托管账户
	err := servicesrpc.AccountRemove(id)
	//TODO: 更新数据库相关信息

	//TODO: 返回删除结果

	return err
}

type ApplyRequest struct {
	Amount int64  `json:"amount"`
	Memo   string `json:"memo"`
}

// ApplyInvoice 使用指定账户申请一张发票
func ApplyInvoice(user *models.User, account *models.Account, applyRequest *ApplyRequest) (*lnrpc.AddInvoiceResponse, error) {
	//获取马卡龙路径
	var macaroonFile string
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir

	if account.UserAccountCode == "admin" {
		macaroonFile = config.GetConfig().ApiConfig.Lnd.MacaroonPath
	} else {
		macaroonFile = filepath.Join(macaroonDir, account.UserAccountCode+".macaroon")
	}
	if macaroonFile == "" {
		CUST.Error("macaroon file not found")
		return nil, fmt.Errorf("macaroon file not found")
	}
	//调用Lit节点发票申请接口
	invoice, err := servicesrpc.InvoiceCreate(applyRequest.Amount, applyRequest.Memo, macaroonFile)
	if err != nil {
		CUST.Error(err.Error())
		return nil, err
	}
	//获取发票信息
	info, _ := FindInvoice(invoice.RHash)

	//构建invoice对象
	var invoiceModel models.Invoice
	invoiceModel.UserID = user.ID
	invoiceModel.Invoice = invoice.PaymentRequest
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
	if err != nil {
		CUST.Error(err.Error())
		return invoice, err
	}
	return invoice, nil
}

type PayInvoiceRequest struct {
	Invoice  string `json:"invoice"`
	FeeLimit int64  `json:"feeLimit"`
}

// PayInvoice 使用指定账户支付发票
func PayInvoice(account *models.Account, PayInvoiceRequest *PayInvoiceRequest) (*lnrpc.Payment, error) {
	//检查数据库中是否有该发票的记录
	a, err := GenericQueryByObject(&models.Balance{
		Invoice: &PayInvoiceRequest.Invoice,
	})
	if err != nil {
		CUST.Error(err.Error())
		return nil, err
	}
	if len(a) > 0 {
		for _, v := range a {
			if v.State == PAY_SUCCESS {
				CUST.Info("该发票已支付")
				return nil, fmt.Errorf("该发票已支付")
			}
			if v.State == PAY_UNKNOWN {
				CUST.Info("该发票支付状态未知")
				return nil, fmt.Errorf("该发票支付状态未知")
			}
		}
	}

	// 判断账户余额是否足够
	info, err := DecodeInvoice(PayInvoiceRequest.Invoice)
	if err != nil {
		CUST.Error("发票解析失败")
		return nil, fmt.Errorf("发票解析失败")
	}

	userBalance, err := QueryCustodyAccount(account.UserAccountCode)
	if err != nil {
		CUST.Error("查询账户余额失败")
		return nil, fmt.Errorf("查询账户余额失败")
	}
	if info.NumSatoshis > userBalance.CurrentBalance {
		CUST.Info("账户余额不足")
		return nil, fmt.Errorf("账户余额不足")
	}

	//获取马卡龙路径
	var macaroonFile string
	macaroonDir := config.GetConfig().ApiConfig.CustodyAccount.MacaroonDir

	if account.UserAccountCode == "admin" {
		macaroonFile = config.GetConfig().ApiConfig.Lnd.MacaroonPath
	} else {
		macaroonFile = filepath.Join(macaroonDir, account.UserAccountCode+".macaroon")
	}
	if macaroonFile == "" {
		CUST.Error("macaroon file not found")
		return nil, fmt.Errorf("macaroon file not found")
	}

	payment, err := servicesrpc.InvoicePay(macaroonFile, PayInvoiceRequest.Invoice, PayInvoiceRequest.FeeLimit)
	if err != nil {
		CUST.Error("pay invoice fail")
		return nil, fmt.Errorf("pay invoice fail")
	}
	var balanceModel models.Balance
	balanceModel.AccountId = account.ID
	balanceModel.BillType = BILL_TYPE_PAYMENT
	balanceModel.Away = AWAY_OUT
	balanceModel.Amount = float64(payment.ValueSat)
	balanceModel.Unit = UNIT_SATOSHIS
	balanceModel.Invoice = &payment.PaymentRequest
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
		CUST.Error(err.Error())
		return payment, err
	}
	return payment, nil
}

// DecodeInvoice  解析发票信息
func DecodeInvoice(invoice string) (*lnrpc.PayReq, error) {
	return servicesrpc.InvoiceDecode(invoice)
}

// FindInvoice 查询节点内部发票
func FindInvoice(rHash []byte) (*lnrpc.Invoice, error) {
	return servicesrpc.InvoiceFind(rHash)
}

// TrackPayment 跟踪支付状态
func TrackPayment(paymentHash string) (*lnrpc.Payment, error) {
	return servicesrpc.PaymentTrack(paymentHash)
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

// PollPayment 遍历所有未确认的发票，轮询支付状态
func pollPayment() {
	//查询数据库，获取所有未确认的支付
	params := QueryParams{
		"State": PAY_UNKNOWN,
	}
	a, err := GenericQuery(&models.Balance{}, params)
	if err != nil {
		CUST.Error(err.Error())
		return
	}
	if len(a) > 0 {
		for _, v := range a {
			temp, err := TrackPayment(*v.PaymentHash)
			if err != nil {
				CUST.Warning(err.Error())
				continue
			}
			if temp.Status == lnrpc.Payment_SUCCEEDED {
				v.State = PAY_SUCCESS
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					CUST.Warning(err.Error())
				}
			} else if temp.Status == lnrpc.Payment_FAILED {
				v.State = PAY_FAILED
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					CUST.Warning(err.Error())
				}
			}
		}
	}
}

// PollInvoice 遍历所有未支付的发票，轮询发票状态
func pollInvoice() {
	//查询数据库，获取所有未支付的发票
	params := QueryParams{
		"Status": lnrpc.Invoice_OPEN,
	}
	a, err := GenericQuery(&models.Invoice{}, params)
	if err != nil {
		CUST.Error(err.Error())
		return
	}
	if len(a) > 0 {
		for _, v := range a {
			invoice, err := DecodeInvoice(v.Invoice)
			if err != nil {
				CUST.Warning(err.Error())
				continue
			}
			rHash, err := hex.DecodeString(invoice.PaymentHash)
			if err != nil {
				CUST.Warning(err.Error())
				continue
			}
			temp, err := FindInvoice(rHash)
			if err != nil {
				CUST.Warning(err.Error())
				continue
			}
			if int16(temp.State) != v.Status {
				v.Status = int16(temp.State)
				mutex.Lock()
				defer mutex.Unlock()
				err = middleware.DB.Save(&v).Error
				if err != nil {
					CUST.Warning(err.Error())
				}
			}
		}
	}
}

func (sm *CronService) PollPaymentCron() {
	CUST.Info("start cron job: PollPayment")
	pollPayment()
}
func (sm *CronService) PollInvoiceCron() {
	CUST.Info("start cron job: PollInvoice")
	pollInvoice()
}
