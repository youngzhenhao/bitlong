package account

import (
	"github.com/btcsuite/btclog"
	"github.com/lightninglabs/pool/account/watcher"
	"github.com/lightningnetwork/lnd/build"
)

const Subsystem = "ACCT"

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log btclog.Logger

// The default amount of logging is none.
func init() {
	UseLogger(build.NewSubLogger(Subsystem, nil))
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using btclog.
func UseLogger(logger btclog.Logger) {
	log = logger

	watcher.UseLogger(logger)
}
