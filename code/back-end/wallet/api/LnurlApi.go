package api

// LnurlUploadUserInfo
//
// @Description: Alice's upload workflow, call after Alice and server's web services are launched
// Alice's front-end uses this LNURL to generate a QR code that waits to be scanned
// @param name
// @param socket
// @return string
// @dev call to get lnurl to generate qr code
func LnurlUploadUserInfo(name, socket string) string {
	return PostServerToUploadUserInfo(name, socket)
}

// LnurlPayToLnu
//
// @Description: Bob's pay-to-lnurl workflow, call after Alice's LNURL QR code is generated
// Bob's front-end scans the Alice's QR code to get the LNURL and then calls the PayToLnurl with amount which Bob wanna pay
// @param ln
// @param amount
// @return string
// @dev call to pay to lnu with amount
func LnurlPayToLnu(lnu, amount string) string {
	invoice := PostServerToPayByPhoneAddInvoice(lnu, amount)
	return SendPaymentSync(invoice)
}
