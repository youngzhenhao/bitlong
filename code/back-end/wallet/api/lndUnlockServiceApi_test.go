package api

import (
	"testing"
)

func TestInitwallet(t *testing.T) {
	Initwallet("cd123321")
}

func TestUnlockwallet(t *testing.T) {
	Unlockwallet("cd123321")
}
