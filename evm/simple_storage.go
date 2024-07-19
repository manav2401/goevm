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
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		accounts: make(map[common.Address]types.StateAccount),
		state:    make(map[common.Address]map[common.Hash]common.Hash),
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
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) SetBalance(address common.Address, balance *uint256.Int) {
	if account, ok := s.accounts[address]; ok {
		account.Balance = balance
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) GetBalance(address common.Address) *uint256.Int {
	if account, ok := s.accounts[address]; ok {
		return account.Balance
	}

	return nil
}

func (s *SimpleStorage) SetNonce(address common.Address, nonce uint64) {
	if account, ok := s.accounts[address]; ok {
		account.Nonce = nonce
		s.accounts[address] = account
	}
}

func (s *SimpleStorage) GetNonce(address common.Address) *uint64 {
	if account, ok := s.accounts[address]; ok {
		return &account.Nonce
	}

	return nil
}

func (s *SimpleStorage) SetState(address common.Address, key common.Hash, value common.Hash) {
	if _, ok := s.state[address]; !ok {
		s.state[address] = make(map[common.Hash]common.Hash)
	}

	s.state[address][key] = value
}

func (s *SimpleStorage) GetState(address common.Address, key common.Hash) common.Hash {
	if state, ok := s.state[address]; ok {
		if value, ok := state[key]; ok {
			return value
		}
	}

	return common.Hash{}
}

func (s *SimpleStorage) Close() {}
