package main

import (
	"goevm/evm"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	// PUSH1 5, PUSH1 6, ADD, STOP
	evm := evm.NewEVM(common.Address{}, common.Address{}, 1, []byte{}, []byte{0x60, 0x5, 0x60, 0x6, 0x1, 0x0}, 10000)
	log.Info("Initialized new evm instance, starting execution")
	evm.Run()
	log.Info("Done execution")
}
