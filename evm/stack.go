package evm

import (
	"github.com/holiman/uint256"
)

// TODO: Decide on how to handle errors during overflow and underflow

const MaxStackSize = 1024

type Stack struct {
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
	s.items = append(s.items, *value)
}

// Pop removes the top item from the stack
func (s *Stack) Pop() uint256.Int {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

// Peek returns the top item from the stack
func (s *Stack) Peek() *uint256.Int {
	return &s.items[len(s.items)-1]
}

func (s *Stack) len() int {
	return len(s.items)
}

func (s *Stack) Dup(n int) {
	s.Push(&s.items[s.len()-n])
}

func (s *Stack) Swap(n int) {
	s.items[s.len()-n], s.items[s.len()-1] = s.items[s.len()-1], s.items[s.len()-n]
}
