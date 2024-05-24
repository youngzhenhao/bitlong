package services

const adminUserId uint = 1

func NewRecharge() {

}

// 托管账户划扣费用
func PayAmountToAdmin(payUserId uint, gasFee, serveFee uint64) (uint, error) {
	id, err := PayAmountInside(payUserId, adminUserId, gasFee, serveFee)
	if err != nil {
		CUST.Error("PayAmountToAdmin failed:%s", err)
		return 0, err
	}
	return id, nil
}
