package api

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/lightninglabs/taproot-assets/fn"
	"github.com/lightninglabs/taproot-assets/tapcfg"
	"github.com/lightningnetwork/lnd/signal"
	"net/http"
	"os"
)

func StartTapRoot() {
	// Hook interceptor for os signals.
	shutdownInterceptor, err := signal.Intercept()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Load the configuration, and parse any command line options. This
	// function will also set up logging properly.
	cfg, cfgLogger, err := tapcfg.LoadConfig(shutdownInterceptor)
	if err != nil {
		var e *flags.Error
		if !errors.As(err, &e) || e.Type != flags.ErrHelp {
			// Print error if not due to help request.
			err = fmt.Errorf("failed to load config: %w", err)
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Help was requested, exit normally.
		os.Exit(0)
	}

	// Enable http profiling server if requested.
	if cfg.Profile != "" {
		go func() {
			profileRedirect := http.RedirectHandler("/debug/pprof",
				http.StatusSeeOther)
			http.Handle("/", profileRedirect)
			cfgLogger.Infof("Pprof listening on %v", cfg.Profile)
			//nolint:gosec
			fmt.Println(http.ListenAndServe(cfg.Profile, nil))
		}()
	}

	// Write cpu profile if requested.
	//if cfg.CPUProfile != "" {
	//	f, err := os.Create(cfg.CPUProfile)
	//	if err != nil {
	//		_, _ = fmt.Fprintln(os.Stderr, err)
	//		os.Exit(1)
	//	}
	//	_ = pprof.StartCPUProfile(f)
	//	defer func(f *os.File) {
	//		err := f.Close()
	//		if err != nil {
	//			fmt.Printf("%s f.Close Error: %v\n", GetTimeNow(), err)
	//		}
	//	}(f)
	//	defer pprof.StopCPUProfile()
	//}

	// This concurrent error queue can be used by every component that can
	// raise runtime errors. Using a queue will prevent us from blocking on
	// sending errors to it, as long as the queue is running.
	errQueue := fn.NewConcurrentQueue[error](fn.DefaultQueueSize)
	errQueue.Start()
	defer errQueue.Stop()

	server, err := tapcfg.CreateServerFromConfig(
		cfg, cfgLogger, shutdownInterceptor, errQueue.ChanIn(),
	)
	if err != nil {
		err := fmt.Errorf("error creating server: %v", err)
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = server.RunUntilShutdown(errQueue.ChanOut())
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
