package main

import (
	"fmt"

	"mse/internal/wallet"
)

func (cli *CLI) createWallet() {
	wallets, _ := wallet.NewWallets(cli.nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(cli.nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
