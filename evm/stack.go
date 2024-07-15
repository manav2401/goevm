package evm

import (
	"github.com/holiman/uint256"

	"sync"
)

// TODO: Decide on how to handle errors during overflow and underflow

const MaxStackSize = 1024

type Stack struct {
	mu    sync.RWMutex
	items []uint256.Int // underlying data
}

// NewStack initializes and returns a new stack instance
func NewStack() *Stack {
	return &Stack{
		items: make([]uint256.Int, 0, MaxStackSize),
	}
}

// Push adds a new item to the stack
func (s *Stack) Push(value *uint256.Int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, *value)
}

// Pop removes the top item from the stack
func (s *Stack) Pop() uint256.Int {
	s.mu.Lock()
	defer s.mu.Unlock()

	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

// Peer returns the top item from the stack
func (s *Stack) Peek() *uint256.Int {
	return &s.items[len(s.items)-1]
}
