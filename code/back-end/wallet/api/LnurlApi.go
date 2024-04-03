package api

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/wallet/base"
	"strconv"
)

// LnurlPhoneWebService
//
// @Description: 1. Run Web [Service]
func LnurlPhoneWebService() {
	RouterRunOnPhone()
}

// LnurlGetPortAvailable
//
// @Description: 2. Get an available port
// @return string
func LnurlGetPortAvailable() string {
	return strconv.Itoa(RequestServerGetPortAvailable(base.QueryConfigByKey("LnurlServerHost")))
}

// LnurlGetNewId
//
// @Description: 3. Get a new id
// @return string
func LnurlGetNewId() string {
	return uuid.New().String()
}

// LnurlFrpcRun
//
// @Description: 4. Forwarding [Service]
// @param id
// @param remotePort
// @return bool
func LnurlFrpcRun(id, remotePort string) {
	if !RequestPostServerIsPortListening(remotePort) {
		FrpcConfig(id, remotePort)
		FrpcRun()
	} else {
		err := errors.New("the server's port you requested is listening, try to get another available remote port")
		fmt.Printf("%s LnurlFrpcRun err:%v\n", GetTimeNow(), err)
	}
}

// LnurlUploadUserInfo
//
// @Description: If 4 success, 5. upload user info
// @note: Alice's upload workflow, call after Alice and server's web services are launched
// Alice's front-end uses this LNURL to generate a QR code that waits to be scannedparam id
// @param name
// @param localPort
// @param remotePort
// @return string
// @dev call to get lnurl to generate qr code
func LnurlUploadUserInfo(id, name, localPort, remotePort string) string {
	return PostServerToUploadUserInfo(id, name, localPort, remotePort)
}

// LnurlPayToLnu
//
// @description: 6. send a payment
// @note: Bob's pay-to-lnurl workflow, call after Alice's LNURL QR code is generated
// Bob's front-end scans the Alice's QR code to get the LNURL and then calls the PayToLnurl with amount which Bob wanna pay
// @param ln
// @param amount
// @return string
// @dev call to pay to lnu with amount
func LnurlPayToLnu(lnu, amount string) string {
	invoice := PostServerToPayByPhoneAddInvoice(lnu, amount)
	return SendPaymentSync(invoice)
}
