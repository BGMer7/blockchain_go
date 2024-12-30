package main

import (
	"fmt"
	"log"

	"mse/internal/blockchain"
	"mse/internal/wallet"
)

func (cli *CLI) createBlockchain(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := blockchain.CreateBlockchain(address, cli.nodeID)
	defer bc.DB.Close()

	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
