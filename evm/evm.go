package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

type Address [20]byte

type EVM struct {
	scope         ScopeContext
	table         JumpTable
	executionOpts *ExecutionOpts
}

type ScopeContext struct {
	stack   *Stack
	memory  *Memory
	storage Storage
}

type ExecutionOpts struct {
	pc         uint64
	contract   common.Address
	sender     common.Address
	value      *uint256.Int
	calldata   []byte
	code       []byte
	gas        uint64
	refund     uint64
	stopFlag   bool
	revertFlag bool
	returnData []byte
}

func newScopeContext() ScopeContext {
	stack := NewStack()
	memory := NewMemory()
	storage := NewSimpleStorage()

	return ScopeContext{
		stack,
		memory,
		storage,
	}
}

func NewExecutionOpts(contract common.Address, sender common.Address, value uint64, calldata []byte, code []byte, gas uint64) *ExecutionOpts {
	return &ExecutionOpts{
		pc:         0,
		contract:   contract,
		sender:     sender,
		value:      uint256.NewInt(value),
		calldata:   calldata,
		code:       code,
		gas:        gas,
		refund:     0,
		stopFlag:   false,
		revertFlag: false,
	}
}

func NewEVM(storage Storage, opts *ExecutionOpts) *EVM {
	sc := newScopeContext()
	table := newInstructionSet()
	return &EVM{
		sc,
		table,
		opts,
	}
}

func (evm *EVM) Run() {
	log.Info("Strating execution in evm", "code", evm.executionOpts.code)
	for {
		opcode := evm.GetOp(evm.executionOpts.pc)
		if op, ok := evm.table[opcode]; ok {
			log.Info("Running", "opcode", opcode, "pc", evm.executionOpts.pc)
			op.execute(evm)
		} else {
			log.Error("Unknown opcode", "opcode", opcode)
		}
		evm.executionOpts.pc++
		if evm.executionOpts.stopFlag || evm.executionOpts.revertFlag {
			log.Info("Stop or revert called", "return data", evm.executionOpts.returnData)
			return
		}
		// decrease gas
	}
}

func (evm *EVM) GetOp(n uint64) OpCode {
	if n < uint64(len(evm.executionOpts.code)) {
		return OpCode(evm.executionOpts.code[n])
	}

	return STOP
}
