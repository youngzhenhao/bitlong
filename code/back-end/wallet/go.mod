module github.com/wallet

go 1.22.2

replace (
	github.com/btcsuite/btcd/btcutil v1.1.5 => ../wallet/lib/btcutil@v1.1.5
	github.com/fatedier/frp => ../wallet/lib/frp@v0.56.0
	// This replace is for
	// https://deps.dev/advisory/OSV/GO-2021-0053?from=%2Fgo%2Fgithub.com%252Fgogo%252Fprotobuf%2Fv1.3.1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/lightninglabs/lightning-terminal v0.12.2-alpha => ../wallet/lib/lightning-terminal@v0.12.2
	github.com/lightninglabs/lnd v0.17.4-beta => ../wallet/lib/lnd@v0.17.4
	github.com/lightninglabs/lndclient v0.17.4-1 => ../wallet/lib/lndclient@v0.17.4
	github.com/ulikunitz/xz => github.com/ulikunitz/xz v0.5.11
	github.com/wallet/api => ../wallet/api
    github.com/lightninglabs/faraday v0.2.13-alpha => ../wallet/lib/faraday@v0.2.13
    github.com/lightninglabs/loop v0.26.6-beta => ../wallet/lib/loop@v0.26.6
    github.com/lightninglabs/pool v0.6.4-beta.0.20231003174306-80d8854a0c4b => ../wallet/lib/pool@v0.6.4
    github.com/lightninglabs/taproot-assets v0.3.3-0.20240315091907-f5ef93e9998a => ../wallet/lib/taproot-assets@v0.3.3
	// We want to format raw bytes as hex instead of base64. The forked version
	// allows us to specify that as an option.
	google.golang.org/protobuf => github.com/lightninglabs/protobuf-go-hex-display v1.30.0-hex-display
)
