package main

import (
	"fmt"
	"goevm/evm"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	fmt.Println("Hello world")
	// PUSH1 5, PUSH1 6, ADD, STOP
	evm := evm.NewEVM(common.Address{}, 1, []byte{}, []byte{0x60, 0x5, 0x60, 0x6, 0x1, 0x0}, 10000)
	fmt.Println("Initialised evm, starting to run")
	evm.Run()
	fmt.Println("Done execution")
}
