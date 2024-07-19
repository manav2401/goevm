package simulation

import (
	"goevm/evm"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"
)

func RunSimpleSimulation() {
	// Create a temporary address for simulation
	sender := common.HexToAddress("0x350fbDe850998AAC40f0b9364b4ACeA665a3d08c")

	// Create a new tracer
	tracer := evm.NewTracer()

	// Create a new storage using tracer
	storage := evm.NewSimpleStorage(tracer)
	defer storage.Close()

	// Create a new account and set some balance
	storage.CreateAccount(sender)
	storage.SetBalance(sender, uint256.NewInt(10000))

	// Arithmetic/Comparision/Logical operations
	var opcodes []evm.OpCode = []evm.OpCode{
		evm.PUSH1, 0x5, // Pushes 5 to stack [0x5]
		evm.PUSH1, 0x6, // Pushes 6 to stack [0x5, 0x6]
		evm.ADD,        // Adds the top two elements of the stack [0xb]
		evm.PUSH1, 0x2, // Pushes 2 to stack [0xb, 0x2]
		evm.MUL,        // Multiplies the top two elements of the stack [0x16]
		evm.PUSH1, 0x5, // Push key to stack [0x16, 0x5]
		evm.GT,         // Greater than [0x0]
		evm.PUSH1, 0x1, // Push key to stack [0x0, 0x1]
		evm.OR, // Bitwise OR [0x1]
	}

	// Environment operations
	opcodes = append(opcodes, []evm.OpCode{
		evm.ADDRESS, // Pushes the address to stack [0x1, address]
		evm.BALANCE, // Pushes the balance to stack [0x1, balance(0x0)]
		evm.POP,     // Pops the top element of the stack [0x1]
	}...)

	// Memory and storage operations
	withWrite := []evm.OpCode{
		// Stack: [0x1] (value)
		evm.PUSH1, 0x0, // Pushes 0 to stack [0x1, 0x0] (offset)
		evm.MSTORE,     // Store value at offset in memory (total length = 32)
		evm.PUSH1, 0x2, // Pushes 2 to stack [0x2] (value)
		evm.PUSH1, 0x20, // Pushes 32 to stack [0x2, 0x20] (offset)
		evm.MSTORE,      // Store second value at offset in memory (total length = 64)
		evm.PUSH1, 0x64, // Pushes 100 to stack [0x64] (value)
		evm.PUSH1, 0x20, // Pushes 32 to stack [0x64, 0x20] (to load value from memory)
		evm.MLOAD,      // Load value from memory at offset [0x64, 0x2] (value)
		evm.SSTORE,     // Store value at key in storage (key = 0x2, value = 0x64)
		evm.PUSH1, 0x2, // Pushes 2 to stack [0x2] (key)
		evm.SLOAD, // Load value from storage at key
		evm.STOP,
	}

	opcodes = append(opcodes, withWrite...)
	code := *(*[]byte)(unsafe.Pointer(&opcodes))

	// Initialise EVM instance
	opts := evm.NewExecutionOpts(common.Address{}, sender, 1, []byte{}, code, 42000)
	evm := evm.NewEVM(storage, opts, tracer)

	log.Info("Initialized new evm instance, starting simple simulation", "len", len(code))

	evm.Run()

	log.Info("Done execution, exiting")
}

func RunRemoteSimulation(path string, contractAddress string) {
	// Create a temporary address for simulation
	sender := common.HexToAddress("0x350fbDe850998AAC40f0b9364b4ACeA665a3d08c")

	// Initialise the contract address
	contract := common.HexToAddress(contractAddress)

	// Create a new tracer
	tracer := evm.NewTracer()

	// Create a new storage using tracer
	storage := evm.NewRemoteStorage(path, tracer)
	defer storage.Close()

	// Arithmetic/Comparision/Logical operations
	var opcodes []evm.OpCode = []evm.OpCode{
		evm.PUSH1, 0x5, // Pushes 5 to stack [0x5]
		evm.PUSH1, 0x6, // Pushes 6 to stack [0x5, 0x6]
		evm.ADD,        // Adds the top two elements of the stack [0xb]
		evm.PUSH1, 0x2, // Pushes 2 to stack [0xb, 0x2]
		evm.MUL,        // Multiplies the top two elements of the stack [0x16]
		evm.PUSH1, 0x5, // Push key to stack [0x16, 0x5]
		evm.GT,         // Greater than [0x0]
		evm.PUSH1, 0x1, // Push key to stack [0x0, 0x1]
		evm.OR, // Bitwise OR [0x1]
	}

	// Environment operations
	opcodes = append(opcodes, []evm.OpCode{
		evm.ADDRESS, // Pushes the address to stack [0x1, address]
		evm.BALANCE, // Pushes the balance to stack [0x1, balance(0x0)]
		evm.POP,     // Pops the top element of the stack [0x1]
	}...)

	// Memory and storage operations
	opcodes = append(opcodes, []evm.OpCode{
		// Stack: [0x1] (value)
		evm.PUSH1, 0x0, // Pushes 0 to stack [0x1, 0x0] (offset)
		evm.MSTORE,     // Store value at offset in memory (total length = 32)
		evm.PUSH1, 0x2, // Pushes 2 to stack [0x2] (value)
		evm.PUSH1, 0x20, // Pushes 32 to stack [0x2, 0x20] (offset)
		evm.MSTORE,     // Store second value at offset in memory (total length = 64)
		evm.PUSH1, 0x0, // Pushes 0 to stack [0x0] (key) (val1 in our test contract)
		evm.SLOAD,      // Load value from storage at key
		evm.PUSH1, 0x1, // Pushes 2 to stack [0x1] (key) (val2 in our test contract)
		evm.SLOAD, // Load value from storage at key
		evm.STOP,  // STOP
	}...)

	// Initialise EVM instance
	code := *(*[]byte)(unsafe.Pointer(&opcodes))
	opts := evm.NewExecutionOpts(contract, sender, 1, []byte{}, code, 42000)
	evm := evm.NewEVM(storage, opts, tracer)

	log.Info("Initialized new evm instance, starting remote simulation", "len", len(code))
	evm.Run()
	log.Info("Done execution, exiting")
}
