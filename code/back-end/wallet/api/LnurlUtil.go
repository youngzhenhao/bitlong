package api

import "github.com/fiatjaf/go-lnurl"

func Encode(url string) string {
	en, _ := lnurl.LNURLEncode(url)
	return en
}
func Decode(lnu string) string {
	de, _ := lnurl.LNURLDecode(lnu)
	return de
}
