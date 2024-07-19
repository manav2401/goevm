package main

import (
	"fmt"
	"goevm/simulation"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

var (
	StorageFlag = &cli.StringFlag{
		Name:  "storage",
		Usage: "Type of storage to use for simulation (simple/remote)",
		Value: "simple",
	}
	Datadir = &cli.StringFlag{
		Name:  "datadir",
		Usage: "Path to use for opening remote geth database",
		Value: "",
	}
	ContractAddressFlag = &cli.StringFlag{
		Name:  "contract-address",
		Usage: "Contract address to be used for remote simulation",
		Value: "",
	}
	simulateCommand = &cli.Command{
		Name:   "simulate",
		Usage:  "Simulate EVM opcodes",
		Action: runSimulator,
		Flags: []cli.Flag{
			StorageFlag,
			Datadir,
			ContractAddressFlag,
		},
	}
)

func initSimulator() *cli.App {
	app := cli.NewApp()
	app.Name = "evm-simulator"
	app.Usage = "Simulate EVM opcodes"
	app.Commands = []*cli.Command{simulateCommand}
	return app
}

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	simulator := initSimulator()
	if err := simulator.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runSimulator(c *cli.Context) error {
	storageType := c.String("storage")
	if storageType == "simple" {
		simulation.RunSimpleSimulation()
		return nil
	}

	if storageType == "remote" {
		contractAddress := c.String("contract-address")
		path := c.String("datadir")
		if contractAddress == "" || path == "" {
			log.Error("Contract address and datadir are required for remote simulation")
			return nil
		}
		simulation.RunRemoteSimulation(path, contractAddress)
		return nil
	}

	log.Error("Invalid simulation type, exiting")
	return nil
}
