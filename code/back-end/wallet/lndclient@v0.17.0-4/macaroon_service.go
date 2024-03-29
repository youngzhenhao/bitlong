package lndclient

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/lightningnetwork/lnd/keychain"
	"github.com/lightningnetwork/lnd/kvdb"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/lightningnetwork/lnd/rpcperms"
	"google.golang.org/grpc"
	"gopkg.in/macaroon-bakery.v2/bakery"
	"gopkg.in/macaroon-bakery.v2/bakery/checkers"
	"gopkg.in/macaroon.v2"
)

const (
	// defaultMacaroonTimeout is the default timeout to be used for general
	// macaroon operation, such as for the macaroon db connection, and for
	// creating a context with a timeout when validating macaroons.
	defaultMacaroonTimeout = 5 * time.Second
)

var (
	// sharedKeyNUMSBytes holds the bytes representing the compressed
	// byte encoding of SharedKeyNUMS. It was generated via a
	// try-and-increment approach using the phrase "Shared Secret" with
	// SHA2-256. The code for the try-and-increment approach can be seen
	// here: https://github.com/lightninglabs/lightning-node-connect/tree/master/mailbox/numsgen
	sharedKeyNUMSBytes, _ = hex.DecodeString(
		"0215b5a3e0ef58b101431e1e513dd017d1018b420bd2e89dcd71f45c031f00469e",
	)

	// SharedKeyNUMS is the public key point that can be use for when we
	// are deriving a shared secret key with LND.
	SharedKeyNUMS, _ = btcec.ParsePubKey(sharedKeyNUMSBytes)

	// SharedKeyLocator is a key locator that can be used for deriving a
	// shared secret key with LND.
	SharedKeyLocator = &keychain.KeyLocator{
		Family: 21,
		Index:  0,
	}
)

// MacaroonService handles the creatation and unlocking of a macaroon DB file.
// It is also a wrapper for a macaroon.Service and uses this to create a
// default macaroon for the caller if stateless mode has not been specified.
type MacaroonService struct {
	cfg *MacaroonServiceConfig

	*macaroons.Service
}

// MacaroonServiceConfig holds configuration values used by the MacaroonService.
type MacaroonServiceConfig struct {
	// RootKeyStorage is an implementation of the main
	// bakery.RootKeyStorage interface. This implementation may also
	// concurrenlty implement the larger macaroons.ExtendedRootKeyStore
	// interface as well.
	RootKeyStore bakery.RootKeyStore

	// MacaroonLocation is the value used for a macaroons' "Location" field.
	MacaroonLocation string

	// MacaroonPath is the path to where macaroons should be stored.
	MacaroonPath string

	// StatelessInit should be set to true if no default macaroons should
	// be created and stored on disk.
	StatelessInit bool

	// Checkers are used to add extra validation checks on macaroon.
	Checkers []macaroons.Checker

	// RequiredPerms defines all method paths and the permissions required
	// when accessing those paths.
	RequiredPerms map[string][]bakery.Op

	// Caveats is a list of caveats that will be added to the default
	// macaroons.
	Caveats []checkers.Caveat

	// DBPassword is the password that will be used to encrypt the macaroon
	// db. If DBPassword is not set, then LndClient, EphemeralKey and
	// KeyLocator must be set instead.
	DBPassword []byte

	// LndClient is an LND client that can be used for any lnd queries.
	// This only needs to be set if DBPassword is not set.
	LndClient *LndServices

	// RPCTimeout is the time after which an RPC call will be canceled if
	// it has not yet completed.
	RPCTimeout time.Duration

	// EphemeralKey is a key that will be used to derive a shared secret
	// with LND. This only needs to be set if DBPassword is not set.
	EphemeralKey *btcec.PublicKey

	// KeyLocator is the locator used to derive a shared secret with LND.
	// This only needs to be set if DBPassword is not set.
	KeyLocator *keychain.KeyLocator
}

// NewMacaroonService checks the config values passed in and creates a
// MacaroonService object accordingly.
func NewMacaroonService(cfg *MacaroonServiceConfig) (*MacaroonService, error) {
	// Validate config.
	if cfg.MacaroonLocation == "" {
		return nil, errors.New("no macaroon location provided")
	}

	if cfg.RPCTimeout == 0 {
		cfg.RPCTimeout = defaultRPCTimeout
	} else if cfg.RPCTimeout < 0 {
		return nil, errors.New("can't have a negative rpc timeout")
	}

	if !cfg.StatelessInit && cfg.MacaroonPath == "" {
		return nil, errors.New("a macaroon path must be given if we " +
			"are not in stateless mode")
	}

	if len(cfg.RequiredPerms) == 0 {
		return nil, errors.New("the required permissions must be set " +
			"and contain elements")
	}

	ms := MacaroonService{
		cfg: cfg,
	}

	if len(cfg.DBPassword) != 0 {
		return &ms, nil
	}

	_, extendedKeyStore := ms.cfg.RootKeyStore.(macaroons.ExtendedRootKeyStore)
	if !extendedKeyStore {
		return &ms, nil
	}

	if cfg.LndClient == nil || cfg.EphemeralKey == nil ||
		cfg.KeyLocator == nil {

		return nil, errors.New("must provide an LndClient, ephemeral " +
			"key and key locator if no DBPassword is provided " +
			"so that a shared key can be derived with LND")
	}

	return &ms, nil
}

// Start starts the macaroon validation service, creates or unlocks the
// macaroon database and, if we are not in stateless mode, creates the default
// macaroon if it doesn't exist yet or regenerates the macaroon if the required
// permissions have changed.
func (ms *MacaroonService) Start() error {
	// Create the macaroon authentication/authorization service.
	service, err := macaroons.NewService(
		ms.cfg.RootKeyStore, ms.cfg.MacaroonLocation, ms.cfg.StatelessInit,
		ms.cfg.Checkers...,
	)
	if err != nil {
		return fmt.Errorf("unable to set up macaroon service: %v", err)
	}
	ms.Service = service

	_, extendedKeyStore := ms.cfg.RootKeyStore.(macaroons.ExtendedRootKeyStore)
	switch {
	// The passed root key store doesn't use the extended interface, so we
	// can skip everything below.
	case !extendedKeyStore:
		break

	case len(ms.cfg.DBPassword) != 0:
		// If a non-empty DB password was provided, then use this
		// directly to try and unlock the db.
		err := ms.CreateUnlock(&ms.cfg.DBPassword)
		if err != nil {
			return fmt.Errorf("unable to unlock macaroon DB: %v",
				err)
		}

	default:
		// If an empty DB password was provided, we want to establish a
		// shared secret with LND which we will use as our DB password.
		ctx, cancel := context.WithTimeout(
			context.Background(), ms.cfg.RPCTimeout,
		)
		defer cancel()

		sharedKey, err := ms.cfg.LndClient.Signer.DeriveSharedKey(
			ctx, ms.cfg.EphemeralKey, ms.cfg.KeyLocator,
		)
		if err != nil {
			return fmt.Errorf("unable to derive a shared "+
				"secret with LND: %v", err)
		}

		// Try to unlock the macaroon store with the shared key.
		dbPassword := sharedKey[:]
		err = ms.CreateUnlock(&dbPassword)
		if err == nil {
			// If the db was successfully unlocked, we can continue.
			break
		}

		log.Infof("Macaroon DB could not be unlocked with the " +
			"derived shared key. Attempting to unlock with " +
			"empty password instead")

		// Otherwise, we will attempt to unlock the db with an empty
		// password. If this succeeds, we will re-encrypt it with
		// the shared key.
		dbPassword = []byte{}
		err = ms.CreateUnlock(&dbPassword)
		if err != nil {
			return fmt.Errorf("unable to unlock macaroon DB: %v",
				err)
		}

		log.Infof("Re-encrypting macaroon DB with derived shared key")

		// Attempt to now re-encrypt the DB with the shared key.
		err = ms.ChangePassword(dbPassword, sharedKey[:])
		if err != nil {
			return fmt.Errorf("unable to change the macaroon "+
				"DB password: %v", err)
		}
	}

	// There are situations in which we don't want a macaroon to be created
	// on disk (for example when running inside LiT stateless integrated
	// mode).
	if ms.cfg.StatelessInit {
		return nil
	}

	// If we are not in stateless mode and a macaroon file does exist, we
	// check that the macaroon matches the required permissions. If not, we
	// will delete the macaroon and create a new one.
	if lnrpc.FileExists(ms.cfg.MacaroonPath) {
		matches, err := ms.macaroonMatchesPermissions()
		if err != nil {
			log.Warnf("An error occurred when attempting to match "+
				"the previous macaroon's permissions with the "+
				"current required permissions. This may occur "+
				"if the previous macaroon file is corrupted. "+
				"The path to the file attempted to be used as "+
				"the previous macaroon is: %s. If that file "+
				"is the correct macaroon file and this error "+
				"happens repeatedly on startup, please remove "+
				"the macaroon file manually and restart "+
				"once again to generate a new macaroon.",
				ms.cfg.MacaroonPath)

			return fmt.Errorf("unable to match the previous "+
				"macaroon's permissions with the required "+
				"permissions: %v", err)
		}

		// In case the old macaroon matches the required permissions,
		// we don't need to create a new macaroon.
		if matches {
			return nil
		}

		// Else if the required permissions have been updated, we delete
		// the old macaroon and create a new one.
		log.Infof("Macaroon at %s does not have all required "+
			"permissions. Deleting it and creating a new "+
			"one", ms.cfg.MacaroonPath)
	}

	// We don't offer the ability to rotate macaroon root keys yet, so just
	// use the default one since the service expects some value to be set.
	idCtx := macaroons.ContextWithRootKeyID(
		context.Background(), macaroons.DefaultRootKeyID,
	)

	// We only generate one default macaroon that contains all existing
	// permissions (equivalent to the admin.macaroon in lnd). Custom
	// macaroons can be created through the bakery RPC.
	mac, err := ms.Oven.NewMacaroon(
		idCtx, bakery.LatestVersion, ms.cfg.Caveats,
		extractPerms(ms.cfg.RequiredPerms)...,
	)
	if err != nil {
		return err
	}

	macBytes, err := mac.M().MarshalBinary()
	if err != nil {
		return err
	}

	err = os.WriteFile(ms.cfg.MacaroonPath, macBytes, 0644)

	return err
}

// Stop cleans up the MacaroonService.
func (ms *MacaroonService) Stop() error {
	var shutdownErr error
	if err := ms.Close(); err != nil {
		log.Errorf("Error closing macaroon service: %v", err)
		shutdownErr = err
	}

	rks := ms.cfg.RootKeyStore
	if eRKS, ok := rks.(macaroons.ExtendedRootKeyStore); ok {
		if err := eRKS.Close(); err != nil {
			log.Errorf("Error closing macaroon DB: %v", err)
			shutdownErr = err
		}
	}

	return shutdownErr
}

// extractPerms creates a deduped list of all the perms in a required perms map.
func extractPerms(requiredPerms map[string][]bakery.Op) []bakery.Op {
	entityActionPairs := make(map[string]map[string]struct{})

	for _, perms := range requiredPerms {
		for _, p := range perms {
			if _, ok := entityActionPairs[p.Entity]; !ok {
				entityActionPairs[p.Entity] = make(
					map[string]struct{},
				)
			}

			entityActionPairs[p.Entity][p.Action] = struct{}{}
		}
	}

	// Dedup the permissions.
	perms := make([]bakery.Op, 0)
	for entity, actions := range entityActionPairs {
		for action := range actions {
			perms = append(perms, bakery.Op{
				Entity: entity,
				Action: action,
			})
		}
	}

	return perms
}

// Interceptors creates gRPC server options with the macaroon security
// interceptors.
func (ms *MacaroonService) Interceptors() (grpc.UnaryServerInterceptor,
	grpc.StreamServerInterceptor, error) {

	interceptor := rpcperms.NewInterceptorChain(log, false, nil)
	err := interceptor.Start()
	if err != nil {
		return nil, nil, err
	}

	interceptor.SetWalletUnlocked()
	interceptor.AddMacaroonService(ms.Service)
	for method, permissions := range ms.cfg.RequiredPerms {
		err := interceptor.AddPermission(method, permissions)
		if err != nil {
			return nil, nil, err
		}
	}

	unaryInterceptor := interceptor.MacaroonUnaryServerInterceptor()
	streamInterceptor := interceptor.MacaroonStreamServerInterceptor()
	return unaryInterceptor, streamInterceptor, nil
}

// NewBoltMacaroonStore returns a new bakery.RootKeyStore, backed by a bolt DB
// instance at the specified location.
func NewBoltMacaroonStore(dbPath, dbFileName string,
	dbTimeout time.Duration) (bakery.RootKeyStore, kvdb.Backend, error) {

	db, err := kvdb.GetBoltBackend(&kvdb.BoltBackendConfig{
		DBPath:     dbPath,
		DBFileName: dbFileName,
		DBTimeout:  dbTimeout,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open macaroon "+
			"db: %w", err)
	}

	rks, err := macaroons.NewRootKeyStorage(db)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open init macaroon "+
			"db: %w", err)
	}

	return rks, db, nil
}

// macaroonMatchesPermissions checks if the macaroon at the cfg.MacaroonPath
// matches the required permissions. It returns true if the macaroon matches the
// required permissions.
func (ms *MacaroonService) macaroonMatchesPermissions() (bool, error) {
	macBytes, err := os.ReadFile(ms.cfg.MacaroonPath)
	if err != nil {
		return false, fmt.Errorf("unable to read macaroon path: %v",
			err)
	}

	// Make sure it actually is a macaroon by parsing it.
	oldMac := &macaroon.Macaroon{}
	if err := oldMac.UnmarshalBinary(macBytes); err != nil {
		return false, fmt.Errorf("unable to decode macaroon: %v", err)
	}

	var (
		authChecker   = ms.Checker.Auth(macaroon.Slice{oldMac})
		requiredPerms = extractPerms(ms.cfg.RequiredPerms)
	)

	ctx, cancel := context.WithTimeout(
		context.Background(), defaultMacaroonTimeout,
	)
	defer cancel()

	_, err = authChecker.Allow(ctx, requiredPerms...)
	if err != nil {
		// If an error is returned here, it's most likely because the
		// old macaroon doesn't match the required permissions. We
		// therefore return false but not the error as this is expected
		// behavior.
		return false, nil
	}

	// If the number of ops in the allowed info is not the same as the
	// number of required permissions, i.e. there are fewer required
	// permissions than allowed ops, then the required permissions have been
	// modified to require fewer permissions than the old macaroon has. We
	// therefore need to regenerate the macaroon.
	allowedInfo, err := authChecker.Allowed(ctx)
	if err != nil {
		return false, err
	}
	if len(allowedInfo.OpIndexes) != len(requiredPerms) {
		return false, nil
	}

	// The old macaroon matches the required permissions.
	return true, nil
}
