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

	accountCreationTrace []interface{}
	storageReadTrace     []interface{}
	storageWriteTrace    []interface{}
}

type StackTrace struct {
	stack *Stack
}

type MemoryTrace struct {
	memory *Memory
}

func NewTracer() *Tracer {
	return &Tracer{
		accountCreationTrace: make([]interface{}, 0),
		storageReadTrace:     make([]interface{}, 0),
		storageWriteTrace:    make([]interface{}, 0),
	}
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
	log.Info("### Opcode Trace", "opcode", t.opcode)
	t.stackTrace.stack.Print("### Stack before")
	scope.stack.Print("### Stack after")

	length := t.memoryTrace.memory.Len()
	t.memoryTrace.memory.Print("### Memory before", length)
	scope.memory.Print("### Memory after", length)

	if len(t.accountCreationTrace) > 0 {
		log.Info("***** Account created", t.accountCreationTrace...)
		t.accountCreationTrace = make([]interface{}, 0)
	}
	if len(t.storageReadTrace) > 0 {
		log.Info("***** Storage read", t.storageReadTrace...)
		t.storageReadTrace = make([]interface{}, 0)
	}
	if len(t.storageWriteTrace) > 0 {
		log.Info("***** Storage write", t.storageWriteTrace...)
		t.storageWriteTrace = make([]interface{}, 0)
	}

	fmt.Println("")
}

func (t *Tracer) CaptureAccountCreation(ctx ...interface{}) {
	t.accountCreationTrace = ctx
}

func (t *Tracer) CaptureStorageReads(ctx ...interface{}) {
	t.storageReadTrace = ctx
}

func (t *Tracer) CaptureStorageWrites(ctx ...interface{}) {
	t.storageWriteTrace = ctx
}
