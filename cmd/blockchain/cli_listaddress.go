package main

import (
	"fmt"
	"log"

	"mse/internal/wallet"
)

func (cli *CLI) listAddresses() {
	wallets, err := wallet.NewWallets(cli.nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
