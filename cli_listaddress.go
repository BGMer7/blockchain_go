package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) []string {
	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}

	return addresses
}
