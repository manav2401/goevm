package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

type Tracer struct {
	stackTrace  StackTrace
	memoryTrace MemoryTrace
	opcode      OpCode // current opcode
}

type StackTrace struct {
	stack *Stack
}

type MemoryTrace struct {
	memory *Memory
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) CaptureTxStart(opts *ExecutionOpts) {
	log.Info("### Starting trace")
	log.Info("Transaction details", "from", opts.sender, "contract", opts.contract, "value", opts.value.Uint64(), "gas", opts.gas)
}

func (t *Tracer) CaptureTxEnd() {
	log.Info("### Ending trace")
}

func (t *Tracer) CaptureOpCodeStart(scope ScopeContext, opcode OpCode) {
	// Capture stack
	stackTrace := StackTrace{
		stack: &Stack{items: make([]uint256.Int, scope.stack.len())},
	}
	copy(stackTrace.stack.items, scope.stack.items)
	t.stackTrace = stackTrace

	// Capture memory
	memoryTrace := MemoryTrace{
		memory: &Memory{data: make([]byte, scope.memory.Len())},
	}
	copy(memoryTrace.memory.data, scope.memory.data)
	t.memoryTrace = memoryTrace

	t.opcode = opcode
}

func (t *Tracer) CaptureOpCodeEnd(scope ScopeContext) {
	fmt.Println("")
	log.Info("### Opcode Trace", "opcode", t.opcode)
	t.stackTrace.stack.Print("### Stack before")
	scope.stack.Print("### Stack after")

	len := t.memoryTrace.memory.Len()
	t.memoryTrace.memory.Print("### Memory before", len)
	scope.memory.Print("### Memory after", len)
	fmt.Println("")
}

func (t *Tracer) CaptureStorageReads() {
}

func (t *Tracer) CaptureStorageWrites() {
}
