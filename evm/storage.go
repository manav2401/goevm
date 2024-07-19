package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// Storage defines base methods that any store needs to implement
type Storage interface {
	IsWriteAllowed() bool

	CreateAccount(common.Address)

	SetBalance(common.Address, *uint256.Int)
	GetBalance(common.Address) *uint256.Int

	SetNonce(common.Address, uint64)
	GetNonce(common.Address) *uint64

	SetState(common.Address, common.Hash, common.Hash)
	GetState(common.Address, common.Hash) common.Hash

	Close()
}
