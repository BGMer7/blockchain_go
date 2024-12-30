package main

import "fmt"

func (cli *CLI) createWallet(nodeID string) string {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	result := fmt.Sprintf("Your new address: %s\n", address)
	fmt.Printf(result)
	return address
}
