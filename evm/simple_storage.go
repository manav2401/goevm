package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

// SimpleStorage is an in-memory store with a simple map underneath
type SimpleStorage struct {
	accounts map[common.Address]types.StateAccount
	state    map[common.Address]map[common.Hash]common.Hash

	tracer *Tracer
}

func NewSimpleStorage(tracer *Tracer) *SimpleStorage {
	return &SimpleStorage{
		accounts: make(map[common.Address]types.StateAccount),
		state:    make(map[common.Address]map[common.Hash]common.Hash),
		tracer:   tracer,
	}
}

func (s *SimpleStorage) IsWriteAllowed() bool {
	return true
}

func (s *SimpleStorage) CreateAccount(address common.Address) {
	if _, ok := s.accounts[address]; !ok {
		account := types.StateAccount{
			Nonce:    0,
			Balance:  uint256.NewInt(0),
			Root:     types.EmptyRootHash,
			CodeHash: []byte{},
		}
		if s.tracer != nil {
			s.tracer.CaptureAccountCreation("address", address, "nonce", account.Nonce, "balance", account.Balance.Uint64(), "root", account.Root, "codeHash", account.CodeHash)
		}
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) SetBalance(address common.Address, balance *uint256.Int) {
	if account, ok := s.accounts[address]; ok {
		if s.tracer != nil {
			s.tracer.CaptureStorageWrites("entity", "balance", "address", address, "old", account.Balance.Uint64(), "new", balance.Uint64())
		}
		account.Balance = balance
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) GetBalance(address common.Address) *uint256.Int {
	var balance *uint256.Int
	if account, ok := s.accounts[address]; ok {
		balance = account.Balance
	}

	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "balance", "address", address, "balance", balance.Uint64())
	}

	return balance
}

func (s *SimpleStorage) SetNonce(address common.Address, nonce uint64) {
	if account, ok := s.accounts[address]; ok {
		if s.tracer != nil {
			s.tracer.CaptureStorageWrites("entity", "nonce", "address", address, "old", account.Nonce, "new", nonce)
		}
		account.Nonce = nonce
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) GetNonce(address common.Address) *uint64 {
	var nonce *uint64
	if account, ok := s.accounts[address]; ok {
		nonce = &account.Nonce
	}

	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "nonce", "address", address, "nonce", *nonce)
	}

	return nonce
}

func (s *SimpleStorage) SetState(address common.Address, key common.Hash, value common.Hash) {
	if _, ok := s.state[address]; !ok {
		s.state[address] = make(map[common.Hash]common.Hash)
	}

	if s.tracer != nil {
		s.tracer.CaptureStorageWrites("entity", "state", "address", address, "key", key, "old", s.state[address][key], "new", value)
	}
	s.state[address][key] = value
}

func (s *SimpleStorage) GetState(address common.Address, key common.Hash) common.Hash {
	val := common.Hash{}
	if state, ok := s.state[address]; ok {
		if value, ok := state[key]; ok {
			val = value
		}
	}

	if s.tracer != nil {
		s.tracer.CaptureStorageReads("entity", "state", "address", address, "key", key, "value", val)
	}

	return val
}

func (s *SimpleStorage) Close() {}
