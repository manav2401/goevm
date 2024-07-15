package evm

type OpCode byte

var (
	STOP OpCode = 0x0

	// Math
	ADD        OpCode = 0x1
	MUL        OpCode = 0x2
	SUB        OpCode = 0x3
	DIV        OpCode = 0x4
	SDIV       OpCode = 0x5
	MOD        OpCode = 0x6
	SMOD       OpCode = 0x7
	ADDMOD     OpCode = 0x8
	MULMOD     OpCode = 0x9
	EXP        OpCode = 0xA
	SIGNEXTEND OpCode = 0xB

	// Comparision
	LT     OpCode = 0x10
	GT     OpCode = 0x11
	SLT    OpCode = 0x12
	SGT    OpCode = 0x13
	EQ     OpCode = 0x14
	ISZERO OpCode = 0x15

	// Logic
	AND OpCode = 0x16
	OR  OpCode = 0x17
	XOR OpCode = 0x18
	NOT OpCode = 0x19

	// Bitops
	BYTE OpCode = 0x1A
	SHL  OpCode = 0x1B
	SHR  OpCode = 0x1C
	SAR  OpCode = 0x1D

	// Misc
	SHA3 OpCode = 0x20

	// Ethereum State
	ADDRESS        OpCode = 0x30
	BALANCE        OpCode = 0x31
	ORIGIN         OpCode = 0x32
	CALLER         OpCode = 0x33
	CALLVALUE      OpCode = 0x34
	CALLDATALOAD   OpCode = 0x35
	CALLDATASIZE   OpCode = 0x36
	CALLDATACOPY   OpCode = 0x37
	CODESIZE       OpCode = 0x38
	CODECOPY       OpCode = 0x39
	GASPRICE       OpCode = 0x3A
	EXTCODESIZE    OpCode = 0x3B
	EXTCODECOPY    OpCode = 0x3C
	RETURNDATASIZE OpCode = 0x3D
	RETURNDATACOPY OpCode = 0x3E
	EXTCODEHASH    OpCode = 0x3F
	BLOCKHASH      OpCode = 0x40
	COINBASE       OpCode = 0x41
	TIMESTAMP      OpCode = 0x42
	NUMBER         OpCode = 0x43
	DIFFICULTY     OpCode = 0x44
	GASLIMIT       OpCode = 0x45
	CHAINID        OpCode = 0x46
	SELFBALANCE    OpCode = 0x47
	BASEFEE        OpCode = 0x48

	POP OpCode = 0x50

	// Memory
	MLOAD   OpCode = 0x51
	MSTORE  OpCode = 0x52
	MSTORE8 OpCode = 0x53

	// Storage
	SLOAD  OpCode = 0x54
	SSTORE OpCode = 0x55

	// Jump
	JUMP     OpCode = 0x56
	JUMPI    OpCode = 0x57
	PC       OpCode = 0x58
	JUMPDEST OpCode = 0x5B

	// Transient storage
	TLOAD  OpCode = 0x5c
	TSTORE OpCode = 0x5d

	// Push
	PUSH1  OpCode = 0x60
	PUSH2  OpCode = 0x61
	PUSH3  OpCode = 0x62
	PUSH4  OpCode = 0x63
	PUSH5  OpCode = 0x64
	PUSH6  OpCode = 0x65
	PUSH7  OpCode = 0x66
	PUSH8  OpCode = 0x67
	PUSH9  OpCode = 0x68
	PUSH10 OpCode = 0x69
	PUSH11 OpCode = 0x6A
	PUSH12 OpCode = 0x6B
	PUSH13 OpCode = 0x6C
	PUSH14 OpCode = 0x6D
	PUSH15 OpCode = 0x6E
	PUSH16 OpCode = 0x6F
	PUSH17 OpCode = 0x70
	PUSH18 OpCode = 0x71
	PUSH19 OpCode = 0x72
	PUSH20 OpCode = 0x73
	PUSH21 OpCode = 0x74
	PUSH22 OpCode = 0x75
	PUSH23 OpCode = 0x76
	PUSH24 OpCode = 0x77
	PUSH25 OpCode = 0x78
	PUSH26 OpCode = 0x79
	PUSH27 OpCode = 0x7A
	PUSH28 OpCode = 0x7B
	PUSH29 OpCode = 0x7C
	PUSH30 OpCode = 0x7D
	PUSH31 OpCode = 0x7E
	PUSH32 OpCode = 0x7F

	// Dup
	DUP1  OpCode = 0x80
	DUP2  OpCode = 0x81
	DUP3  OpCode = 0x82
	DUP4  OpCode = 0x83
	DUP5  OpCode = 0x84
	DUP6  OpCode = 0x85
	DUP7  OpCode = 0x86
	DUP8  OpCode = 0x87
	DUP9  OpCode = 0x88
	DUP10 OpCode = 0x89
	DUP11 OpCode = 0x8A
	DUP12 OpCode = 0x8B
	DUP13 OpCode = 0x8C
	DUP14 OpCode = 0x8D
	DUP15 OpCode = 0x8E
	DUP16 OpCode = 0x8F

	// Swap
	SWAP1  OpCode = 0x90
	SWAP2  OpCode = 0x91
	SWAP3  OpCode = 0x92
	SWAP4  OpCode = 0x93
	SWAP5  OpCode = 0x94
	SWAP6  OpCode = 0x95
	SWAP7  OpCode = 0x96
	SWAP8  OpCode = 0x97
	SWAP9  OpCode = 0x98
	SWAP10 OpCode = 0x99
	SWAP11 OpCode = 0x9A
	SWAP12 OpCode = 0x9B
	SWAP13 OpCode = 0x9C
	SWAP14 OpCode = 0x9D
	SWAP15 OpCode = 0x9E
	SWAP16 OpCode = 0x9F

	// Log
	LOG0 OpCode = 0xA0
	LOG1 OpCode = 0xA1
	LOG2 OpCode = 0xA2
	LOG3 OpCode = 0xA3
	LOG4 OpCode = 0xA4

	// Contract
	CREATE       OpCode = 0xF0
	CALL         OpCode = 0xF1
	CALLCODE     OpCode = 0xF2
	RETURN       OpCode = 0xF3
	DELEGATECALL OpCode = 0xF4
	CREATE2      OpCode = 0xF5
	STATICCALL   OpCode = 0xFA
	REVERT       OpCode = 0xFD
	INVALID      OpCode = 0xFE
	SELFDESTRUCT OpCode = 0xFF
)
