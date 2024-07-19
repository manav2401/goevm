package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

// RemoteStorage represents a disk based store of an existing geth based
// EVM node (hash based scheme). It acts as an interafce to interact with
// the underlying data (e.g. state and accounts) from the node.
type RemoteStorage struct {
	root    common.Hash    // latest head's root hash
	db      ethdb.Database // for raw kv interactions
	statedb state.Database // for accessing storage tries whenever required
	trie    state.Trie     // for accessing main merkle trie

	tracer *Tracer
}

func NewRemoteStorage(path string, tracer *Tracer) *RemoteStorage {
	// Open the key value db given the path. We're assuming that it's using leveldb
	db, err := rawdb.NewLevelDBDatabase(path, 1024, 2000, "", true)
	if err != nil {
		log.Error("Error opening leveldb database", "path", path, "err", err)
		return nil
	}

	// Open the trie database with the kv database and default hash scheme based config
	trieDb := triedb.NewDatabase(db, triedb.HashDefaults)
	if trieDb == nil {
		log.Error("Error opening trie database", "path", path)
		db.Close()
		return nil
	}

	// Create a state database
	stateDb := state.NewDatabaseWithNodeDB(db, trieDb)

	// Get the latest head from db
	latest := rawdb.ReadHeadHeader(db)
	if latest == nil {
		log.Error("Unable to query latest header from kv db")
		db.Close()
		return nil
	}

	// Open the trie using the latest head's root
	trie, err := stateDb.OpenTrie(latest.Root)
	if err != nil {
		log.Error("Unable to open trie on latest head's root", "root", latest.Root, "number", latest.Number.Uint64(), "hash", latest.Hash(), "err", err)
		db.Close()
		return nil
	}

	log.Info("Opened database using latest head", "number", latest.Number.Uint64(), "root", latest.Root, "hash", latest.Hash())

	return &RemoteStorage{
		root:    latest.Root,
		db:      db,
		statedb: stateDb,
		trie:    trie,
		tracer:  tracer,
	}
}

func (s *RemoteStorage) IsWriteAllowed() bool {
	return false
}

func (s *RemoteStorage) CreateAccount(common.Address) {}

func (s *RemoteStorage) SetBalance(common.Address, *uint256.Int) {}

func (s *RemoteStorage) GetBalance(address common.Address) *uint256.Int {
	account, err := s.trie.GetAccount(address)
	if err != nil {
		log.Error("Error getting account from db", "address", address, "err", err)
		return nil
	}
	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "balance", "address", address, "balance", account.Balance.Uint64())
	}
	return account.Balance
}

func (s *RemoteStorage) SetNonce(common.Address, uint64) {}

func (s *RemoteStorage) GetNonce(address common.Address) *uint64 {
	account, err := s.trie.GetAccount(address)
	if err != nil {
		log.Error("Error getting account from db", "address", address, "err", err)
		return nil
	}
	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "nonce", "address", address, "nonce", account.Nonce)
	}
	return &account.Nonce
}

func (s *RemoteStorage) SetState(common.Address, common.Hash, common.Hash) {}

func (s *RemoteStorage) GetState(address common.Address, key common.Hash) common.Hash {
	storageTrie := openStorageTrie(address, s.root, s.trie, s.statedb)
	if storageTrie == nil {
		return common.Hash{}
	}

	val, err := storageTrie.GetStorage(address, key.Bytes())
	if err != nil {
		log.Error("Error getting data from storage trie", "address", address, "key", key, "err", err)
		return common.Hash{}
	}

	var value common.Hash
	value.SetBytes(val)
	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "state", "address", address, "key", key, "value", value)
	}
	return value
}

func openStorageTrie(address common.Address, root common.Hash, globalTrie state.Trie, statedb state.Database) state.Trie {
	account, err := globalTrie.GetAccount(address)
	if err != nil {
		log.Error("Error getting account from db", "address", address, "err", err)
		return nil
	}

	// Open the storage trie for the given contract address
	trie, err := statedb.OpenStorageTrie(root, address, account.Root, globalTrie)
	if err != nil {
		log.Error("Error opening storage trie", "address", address, "root", account.Root, "err", err)
		return nil
	}

	return trie
}

func (s *RemoteStorage) Close() {
	s.db.Close()
}
