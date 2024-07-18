package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// Storage defines base methods that any store needs to implement
type Storage interface {
	GetBalance(common.Address) *uint256.Int
	GetNonce(common.Address) *uint64
}
