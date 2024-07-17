package evm

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

func opStop(evm *EVM) ([]byte, error) {
	evm.executionOpts.stopFlag = true
	return nil, nil
}

func opAdd(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Add(&x, y)
	return nil, nil
}

func opMul(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Mul(&x, y)
	return nil, nil
}

func opSub(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Sub(&x, y)
	return nil, nil
}

func opDiv(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Div(&x, y)
	return nil, nil
}

func opSDiv(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.SDiv(&x, y)
	return nil, nil
}

func opMod(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Mod(&x, y)
	return nil, nil
}

func opSMod(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.SMod(&x, y)
	return nil, nil
}

func opAddMod(evm *EVM) ([]byte, error) {
	x, y, z := evm.scope.stack.Pop(), evm.scope.stack.Pop(), evm.scope.stack.Peek()
	z.AddMod(&x, &y, z)
	return nil, nil
}

func opMulMod(evm *EVM) ([]byte, error) {
	x, y, z := evm.scope.stack.Pop(), evm.scope.stack.Pop(), evm.scope.stack.Peek()
	z.MulMod(&x, &y, z)
	return nil, nil
}

func opExp(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Exp(&x, y)
	return nil, nil
}

func opSignExtend(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.ExtendSign(&x, y)
	return nil, nil
}

func opLt(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if x.Lt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

func opGt(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if x.Gt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

func opSlt(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if x.Slt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

func opSgt(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if x.Sgt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

func opEq(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if x.Eq(y) {
		y.SetOne()
	} else {
		y.Clear()
	}
	return nil, nil
}

func opIsZero(evm *EVM) ([]byte, error) {
	x := evm.scope.stack.Peek()
	if x.IsZero() {
		x.SetOne()
	} else {
		x.Clear()
	}
	return nil, nil
}

func opAnd(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.And(&x, y)
	return nil, nil
}

func opOr(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Or(&x, y)
	return nil, nil
}

func opXor(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Xor(&x, y)
	return nil, nil
}

func opNot(evm *EVM) ([]byte, error) {
	x := evm.scope.stack.Peek()
	x.Not(x)
	return nil, nil
}

func opByte(evm *EVM) ([]byte, error) {
	x, y := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	y.Byte(&x)
	return nil, nil
}

func opShl(evm *EVM) ([]byte, error) {
	shift, value := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if shift.LtUint64(256) {
		value.Lsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}
	return nil, nil
}

func opShr(evm *EVM) ([]byte, error) {
	shift, value := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if shift.LtUint64(256) {
		value.Rsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}
	return nil, nil
}

func opSar(evm *EVM) ([]byte, error) {
	shift, value := evm.scope.stack.Pop(), evm.scope.stack.Peek()
	if shift.GtUint64(256) {
		if value.Sign() >= 0 {
			value.Clear()
		} else {
			// Max negative shift: all bits set
			value.SetAllOne()
		}
		return nil, nil
	}
	n := uint(shift.Uint64())
	value.SRsh(value, n)
	return nil, nil
}

func opAddress(evm *EVM) ([]byte, error) {
	evm.scope.stack.Push(new(uint256.Int).SetBytes(evm.executionOpts.sender.Bytes()))
	return nil, nil
}

func opBalance(evm *EVM) ([]byte, error) {
	slot := evm.scope.stack.Peek()
	address := common.Address(slot.Bytes20())
	slot.Set(evm.scope.storage.GetBalance(address))
	return nil, nil
}

func opCaller(evm *EVM) ([]byte, error) {
	// For simplicity, we set the tx sender on stack
	evm.scope.stack.Push(new(uint256.Int).SetBytes(evm.executionOpts.sender.Bytes()))
	return nil, nil
}

func opCallValue(evm *EVM) ([]byte, error) {
	evm.scope.stack.Push(evm.executionOpts.value)
	return nil, nil
}

func opCalldataLoad(evm *EVM) ([]byte, error) {
	x := evm.scope.stack.Peek()
	if offset, overflow := x.Uint64WithOverflow(); !overflow {
		data := getData(evm.executionOpts.calldata, offset, 32)
		x.SetBytes(data)
	} else {
		x.Clear()
	}
	return nil, nil
}

func opCalldataSize(evm *EVM) ([]byte, error) {
	evm.scope.stack.Push(new(uint256.Int).SetUint64(uint64(len(evm.executionOpts.calldata))))
	return nil, nil
}

func opCalldataCopy(evm *EVM) ([]byte, error) {
	var (
		memOffset  = evm.scope.stack.Pop()
		dataOffset = evm.scope.stack.Pop()
		length     = evm.scope.stack.Pop()
	)
	dataOffset64, overflow := dataOffset.Uint64WithOverflow()
	if overflow {
		dataOffset64 = math.MaxUint64
	}
	// These values are checked for overflow during gas cost calculation
	memOffset64 := memOffset.Uint64()
	length64 := length.Uint64()

	// TODO: think a better way/place to resize
	evm.scope.memory.Resize(evm.scope.memory.Len() + length64)
	evm.scope.memory.Store(memOffset64, length64, getData(evm.executionOpts.calldata, dataOffset64, length64))
	return nil, nil
}

func opCodesize(evm *EVM) ([]byte, error) {
	evm.scope.stack.Push(new(uint256.Int).SetUint64(uint64(len(evm.executionOpts.code))))
	return nil, nil
}

func opCodeCopy(evm *EVM) ([]byte, error) {
	var (
		memOffset  = evm.scope.stack.Pop()
		codeOffset = evm.scope.stack.Pop()
		length     = evm.scope.stack.Pop()
	)
	uint64CodeOffset, overflow := codeOffset.Uint64WithOverflow()
	if overflow {
		uint64CodeOffset = math.MaxUint64
	}

	codeCopy := getData(evm.executionOpts.code, uint64CodeOffset, length.Uint64())

	// TODO: think a better way/place to resize
	evm.scope.memory.Resize(evm.scope.memory.Len() + length.Uint64())
	evm.scope.memory.Store(memOffset.Uint64(), length.Uint64(), codeCopy)
	return nil, nil
}

func opPop(evm *EVM) ([]byte, error) {
	evm.scope.stack.Pop()
	return nil, nil
}

func makePush(size uint64) executeFn {
	return func(evm *EVM) ([]byte, error) {
		if size == 0 {
			evm.scope.stack.Push(new(uint256.Int))
			return nil, nil
		}
		start := evm.executionOpts.pc + 1
		end := start + size
		if int(end) > len(evm.executionOpts.code) {
			panic("not enough bytes in code to push")
		}
		evm.scope.stack.Push(new(uint256.Int).SetBytes(
			common.RightPadBytes(evm.executionOpts.code[start:end], int(size)),
		))

		// This will bring the pc to last byte (increment for next opcode won't happen here)
		evm.executionOpts.pc += size
		return nil, nil
	}
}

func opMload(evm *EVM) ([]byte, error) {
	value := evm.scope.stack.Peek()
	offset := value.Uint64()
	value.SetBytes(evm.scope.memory.Load(offset, 32))
	return nil, nil
}

func opMstore(evm *EVM) ([]byte, error) {
	offset, value := evm.scope.stack.Pop(), evm.scope.stack.Pop()
	valueB32 := value.Bytes32()

	// TODO: think a better way/place to resize
	evm.scope.memory.Resize(evm.scope.memory.Len() + 32)
	evm.scope.memory.Store(offset.Uint64(), 32, valueB32[:])
	return nil, nil
}

func opMstore8(evm *EVM) ([]byte, error) {
	offset, value := evm.scope.stack.Pop(), evm.scope.stack.Pop()
	v := []byte{byte(value.Uint64())}

	// TODO: think a better way/place to resize
	evm.scope.memory.Resize(evm.scope.memory.Len() + 1)
	evm.scope.memory.Store(offset.Uint64(), 1, v)
	return nil, nil
}

func opSload(evm *EVM) ([]byte, error) {
	evm.scope.stack.Pop()
	return nil, nil
}

func opSStore(evm *EVM) ([]byte, error) {
	evm.scope.stack.Pop()
	return nil, nil
}
