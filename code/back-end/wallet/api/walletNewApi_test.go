package api

import (
	"testing"
)

func TestGetAddr(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			Mnemonic     string
			Account      int64
			addressIndex int64
		}
	}{
		{
			name: "Test Case 1",
			args: struct {
				Mnemonic     string
				Account      int64
				addressIndex int64
			}{
				Mnemonic:     "exhaust answer inject buddy enrich peasant frozen fish birth grape speak yellow cactus hundred mad fiber surprise essay lock maximum share area heavy final",
				Account:      0,
				addressIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAddr(tt.args.Mnemonic, tt.args.Account, tt.args.addressIndex)
		})
	}
}
