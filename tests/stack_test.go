package tests

import (
	"goevm/evm"
	"testing"

	"github.com/holiman/uint256"
)

func TestStack(t *testing.T) {
	stack := evm.NewStack()
	one := uint256.NewInt(1)
	two := uint256.NewInt(1)
	stack.Push(one)
	stack.Push(two)
	top := stack.Peek()
	if top.Cmp(two) != 0 {
		t.Fatalf("Invalid peer, expected: %v and got: %v", two.Uint64(), top.Uint64())
	}
	*top = stack.Pop()
	if top.Cmp(two) != 0 {
		t.Fatalf("Invalid pop, expected: %v and got: %v", two.Uint64(), top.Uint64())
	}
	top = stack.Peek()
	if top.Cmp(one) != 0 {
		t.Fatalf("Invalid peek, expected: %v and got: %v", one.Uint64(), top.Uint64())
	}
	*top = stack.Pop()
	if top.Cmp(one) != 0 {
		t.Fatalf("Invalid pop, expected: %v and got: %v", one.Uint64(), top.Uint64())
	}
}
