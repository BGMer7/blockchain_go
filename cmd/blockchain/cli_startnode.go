package main

import (
	"fmt"
	"log"

	"mse/internal/network"
	"mse/internal/wallet"
)

func (cli *CLI) startNode(minerAddress string) {
	fmt.Printf("Starting node %s\n", cli.nodeID)
	if len(minerAddress) > 0 {
		if !wallet.ValidateAddress(minerAddress) {
			log.Panic("Wrong miner address!")
		}
		fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
	}
	network.StartServer(cli.nodeID, minerAddress)
}
