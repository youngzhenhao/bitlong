package api

import (
	"testing"
)

func TestInitwallet(t *testing.T) {
	InitWallet("12345678")
}

func TestUnlockwallet(t *testing.T) {
	UnlockWallet("12345678")
}
