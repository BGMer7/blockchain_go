package main

import (
	"fmt"

	"mse/internal/blockchain"
)

func (cli *CLI) reindexUTXO() {
	bc := blockchain.NewBlockchain(cli.nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
