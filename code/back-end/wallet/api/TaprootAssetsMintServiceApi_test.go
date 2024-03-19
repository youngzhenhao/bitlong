package api

import (
	"testing"
)

func TestInitMint(t *testing.T) {
	InitMint()
}
func TestFinalizeMint(t *testing.T) {
	FinalizeMint()
}
func TestGetTapRootAddr(t *testing.T) {
	GetTapRootAddr("", 10)
}

func TestSendAssets(t *testing.T) {
	SendAssets()

}
