package evm

type JumpTable map[OpCode]OpCodeOperation

type executeFn func(*EVM) ([]byte, error)

type OpCodeOperation struct {
	gas     uint64
	execute executeFn
}

func newInstructionSet() JumpTable {
	table := make(map[OpCode]OpCodeOperation)

	table[STOP] = OpCodeOperation{0, opStop}

	table[ADD] = OpCodeOperation{3, opAdd}
	table[MUL] = OpCodeOperation{5, opMul}
	table[SUB] = OpCodeOperation{3, opSub}
	table[DIV] = OpCodeOperation{5, opDiv}
	table[SDIV] = OpCodeOperation{5, opSDiv}
	table[MOD] = OpCodeOperation{5, opMod}
	table[SMOD] = OpCodeOperation{5, opSMod}
	table[ADDMOD] = OpCodeOperation{8, opAddMod}
	table[MULMOD] = OpCodeOperation{8, opMulMod}
	table[EXP] = OpCodeOperation{10, opExp}
	table[SIGNEXTEND] = OpCodeOperation{5, opSignExtend}

	table[LT] = OpCodeOperation{3, opLt}
	table[GT] = OpCodeOperation{3, opGt}
	table[SLT] = OpCodeOperation{3, opSlt}
	table[SGT] = OpCodeOperation{3, opSgt}
	table[EQ] = OpCodeOperation{3, opEq}
	table[ISZERO] = OpCodeOperation{3, opIsZero}

	table[AND] = OpCodeOperation{3, opAnd}
	table[OR] = OpCodeOperation{3, opOr}
	table[XOR] = OpCodeOperation{3, opXor}
	table[NOT] = OpCodeOperation{3, opNot}
	table[BYTE] = OpCodeOperation{3, opByte}
	table[SHL] = OpCodeOperation{3, opShl}
	table[SHR] = OpCodeOperation{3, opShr}
	table[SAR] = OpCodeOperation{3, opSar}

	table[ADDRESS] = OpCodeOperation{2, opAddress}
	table[BALANCE] = OpCodeOperation{100, opBalance} // only warm considered
	table[ORIGIN] = OpCodeOperation{2, opOrigin}
	table[CALLER] = OpCodeOperation{2, opCaller}
	table[CALLVALUE] = OpCodeOperation{2, opCallValue}
	table[CALLDATALOAD] = OpCodeOperation{3, opCalldataLoad}
	table[CALLDATASIZE] = OpCodeOperation{2, opCalldataSize}
	table[CALLDATACOPY] = OpCodeOperation{3, opCalldataCopy} // only static considered
	table[CODESIZE] = OpCodeOperation{2, opCodesize}
	table[CODECOPY] = OpCodeOperation{3, opCodeCopy} // only static considered

	table[POP] = OpCodeOperation{2, opPop}
	table[PUSH0] = OpCodeOperation{2, makePush(0)}
	for i := 0; i < 32; i++ {
		op := PUSH1 + OpCode(i)
		table[op] = OpCodeOperation{3, makePush(uint64(i + 1))}
	}

	table[MLOAD] = OpCodeOperation{3, opMload}     // only static considered
	table[MSTORE] = OpCodeOperation{3, opMstore}   // only static considered
	table[MSTORE8] = OpCodeOperation{3, opMstore8} // only static considered
	table[SLOAD] = OpCodeOperation{100, opSload}   // only warm considered
	table[SSTORE] = OpCodeOperation{100, opSStore} // only warm considered

	table[JUMP] = OpCodeOperation{8, opJump}
	table[JUMPI] = OpCodeOperation{10, opJumpi}
	table[PC] = OpCodeOperation{2, opPc}
	table[JUMPDEST] = OpCodeOperation{1, opJumpdest}

	for i := 0; i < 16; i++ {
		op := DUP1 + OpCode(i)
		table[op] = OpCodeOperation{3, makeDup(i + 1)}
	}

	for i := 0; i < 16; i++ {
		op := SWAP1 + OpCode(i)
		table[op] = OpCodeOperation{3, makeSwap(i + 1)}
	}

	table[RETURN] = OpCodeOperation{0, opReturn}
	table[REVERT] = OpCodeOperation{0, opRevert}

	return table
}
