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

	table[ADD] = OpCodeOperation{0, opAdd}
	table[MUL] = OpCodeOperation{0, opMul}
	table[SUB] = OpCodeOperation{0, opSub}
	table[DIV] = OpCodeOperation{0, opDiv}
	table[SDIV] = OpCodeOperation{0, opSDiv}
	table[MOD] = OpCodeOperation{0, opMod}
	table[SMOD] = OpCodeOperation{0, opSMod}
	table[ADDMOD] = OpCodeOperation{0, opAddMod}
	table[MULMOD] = OpCodeOperation{0, opMulMod}
	table[EXP] = OpCodeOperation{0, opExp}
	table[SIGNEXTEND] = OpCodeOperation{0, opSignExtend}

	table[LT] = OpCodeOperation{0, opLt}
	table[GT] = OpCodeOperation{0, opGt}
	table[SLT] = OpCodeOperation{0, opSlt}
	table[SGT] = OpCodeOperation{0, opSgt}
	table[EQ] = OpCodeOperation{0, opEq}
	table[ISZERO] = OpCodeOperation{0, opIsZero}

	table[AND] = OpCodeOperation{0, opAnd}
	table[OR] = OpCodeOperation{0, opOr}
	table[XOR] = OpCodeOperation{0, opXor}
	table[NOT] = OpCodeOperation{0, opNot}
	table[BYTE] = OpCodeOperation{0, opByte}
	table[SHL] = OpCodeOperation{0, opShl}
	table[SHR] = OpCodeOperation{0, opShr}
	table[SAR] = OpCodeOperation{0, opSar}

	table[ADDRESS] = OpCodeOperation{0, opAddress}
	table[BALANCE] = OpCodeOperation{0, opBalance}
	table[CALLER] = OpCodeOperation{0, opCaller}
	table[CALLVALUE] = OpCodeOperation{0, opCallValue}
	table[CALLDATALOAD] = OpCodeOperation{0, opCalldataLoad}
	table[CALLDATASIZE] = OpCodeOperation{0, opCalldataSize}
	table[CALLDATACOPY] = OpCodeOperation{0, opCalldataCopy}
	table[CODESIZE] = OpCodeOperation{0, opCodesize}
	table[CODECOPY] = OpCodeOperation{0, opCodeCopy}

	table[POP] = OpCodeOperation{0, opPop}
	table[PUSH0] = OpCodeOperation{0, makePush(0)}
	for i := 0; i < 32; i++ {
		op := PUSH1 + OpCode(i)
		table[op] = OpCodeOperation{0, makePush(uint64(i + 1))}
	}

	return table
}
