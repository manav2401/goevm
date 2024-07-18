package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// SimpleStorage is an in-memory store with a simple map underneath
type SimpleStorage struct {
	accounts map[common.Address]Account
}

type Account struct {
	Nonce    uint64
	Balance  *uint256.Int
	Root     common.Hash
	CodeHash []byte
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{accounts: make(map[common.Address]Account)}
}

func (s *SimpleStorage) GetBalance(address common.Address) *uint256.Int {
	if account, ok := s.accounts[address]; ok {
		return account.Balance
	}

	return nil
}

func (s *SimpleStorage) GetNonce(address common.Address) *uint64 {
	if account, ok := s.accounts[address]; ok {
		return &account.Nonce
	}

	return nil
}
