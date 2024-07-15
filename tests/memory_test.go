package tests

import (
	"goevm/evm"
	"testing"
)

func TestMemory(t *testing.T) {
	memory := evm.NewMemory()
	memory.Resize(10)
	data := "hello"
	memory.Store(0, 5, []byte(data))
	data = "world"
	memory.Store(5, 5, []byte(data))
	len := memory.Len()
	if len != 10 {
		t.Fatalf("Invalid memory length, expected: %d, got: %d", 5, len)
	}
	storedData := memory.Load(0, 10)
	if string(storedData) != "helloworld" {
		t.Fatalf("Invalid stored data in memory, expected: %s, got: %s", "helloworld", string(storedData))
	}
}
