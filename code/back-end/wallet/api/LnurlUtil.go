package api

import (
	"fmt"
	"github.com/fiatjaf/go-lnurl"
	"net"
)

func Encode(url string) string {
	en, _ := lnurl.LNURLEncode(url)
	return en
}
func Decode(lnu string) string {
	de, _ := lnurl.LNURLDecode(lnu)
	return de
}

// QueryAvailablePort
//
// @note: Query for an available port on this host.
// @dev: Query for unused ports in the port range [1024:49151],
// return an available port
// @return: uint16
func QueryAvailablePort() uint16 {
	var startPort uint16 = 1024
	var endPort uint16 = 49151
	for port := startPort; port <= endPort; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			_ = listener.Close()
			return port
		}
	}
	return 0
}
