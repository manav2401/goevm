package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/holiman/uint256"
)

// RemoteStorage represents a disk based store of an existing geth based
// EVM node (hash based scheme). It acts as an interafce to interact with
// the underlying data (e.g. state and accounts) from the node.
type RemoteStorage struct {
	db   ethdb.Database // for raw kv interactions
	trie state.Trie     // for state interactions
}

func NewRemoteStorage(path string) *RemoteStorage {
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

	return &RemoteStorage{
		db:   db,
		trie: trie,
	}
}

func (s *RemoteStorage) CreateAccount(common.Address, types.StateAccount) {}

func (s *RemoteStorage) SetBalance(common.Address, *uint256.Int) {}

func (s *RemoteStorage) GetBalance(address common.Address) *uint256.Int {
	account, err := s.trie.GetAccount(address)
	if err != nil {
		log.Error("Error getting account from db", "address", address, "err", err)
		return nil
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
	return &account.Nonce
}

func (s *RemoteStorage) SetState(common.Address, common.Hash) {}

func (s *RemoteStorage) GetState(address common.Address, key common.Hash) common.Hash {
	val, err := s.trie.GetStorage(address, key.Bytes())
	if err != nil {
		log.Error("Error getting storage from db", "address", address, "key", key, "err", err)
		return types.EmptyRootHash
	}

	var value common.Hash
	value.SetBytes(val)
	return value
}

func (s *RemoteStorage) Close() {
	s.db.Close()
}
