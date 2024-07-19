package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

const IntrinsicGasCost = 21000

type EVM struct {
	scope         ScopeContext
	table         JumpTable
	executionOpts *ExecutionOpts
	tracer        *Tracer
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

	return ScopeContext{
		stack:  stack,
		memory: memory,
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

func NewEVM(storage Storage, opts *ExecutionOpts, tracer *Tracer) *EVM {
	sc := newScopeContext()
	sc.storage = storage

	table := newInstructionSet()
	return &EVM{
		sc,
		table,
		opts,
		tracer,
	}
}

func (evm *EVM) Run() {
	log.Info("Starting execution in evm")
	if evm.tracer != nil {
		evm.tracer.CaptureTxStart(evm.executionOpts)
		defer evm.tracer.CaptureTxEnd(evm.executionOpts)
	}

	// Check for the intrinsic gas cost and deduct it
	if evm.executionOpts.gas < IntrinsicGasCost {
		log.Error("Insufficient gas to run the code", "gas", evm.executionOpts.gas)
		return
	}
	evm.executionOpts.gas -= IntrinsicGasCost

	for {
		opcode := evm.GetOp(evm.executionOpts.pc)
		if op, ok := evm.table[opcode]; ok {
			cost := op.gas
			if evm.executionOpts.gas < cost {
				log.Error("Insufficient gas to run the opcode", "opcode", opcode, "remaining", evm.executionOpts.gas, "required", cost)
				return
			}
			if evm.tracer != nil {
				evm.tracer.CaptureOpCodeStart(evm.scope, opcode, evm.executionOpts.gas)
			}

			// Capture memory length before executing the opcode
			memLength := evm.scope.memory.Len()

			// Call the execute function of the opcode
			op.execute(evm)

			// Calculate memory expansion cost (3 per byte)
			memCost := (evm.scope.memory.Len() - memLength) * 3
			evm.executionOpts.gas -= cost + memCost

			if evm.tracer != nil {
				evm.tracer.CaptureOpCodeEnd(evm.scope, evm.executionOpts.gas)
			}
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
