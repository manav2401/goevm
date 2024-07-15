package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Key [32]byte
type Value [32]byte

// Storage defines base methods that any store needs to implement
type Storage interface {
	// Store writes the given value against the key to underlying storage
	Store(Key, Value)

	// Load reads the value correponding to the key from underlying storage
	Load(Key) Value

	GetBalance(common.Address) *uint256.Int
}
