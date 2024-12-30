package main

import (
	"fmt"
	"log"

	"mse/internal/blockchain"
	"mse/internal/wallet"
)

func (cli *CLI) send(from, to string, amount int, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(cli.nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	defer bc.DB.Close()

	wallets, err := wallet.NewWallets(cli.nodeID)
	if err != nil {
		log.Panic(err)
	}

	wallet := wallets.GetWallet(from)
	if wallet == nil {
		log.Panic("ERROR: Sender wallet not found")
	}

	tx := blockchain.NewUTXOTransaction(wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := blockchain.NewCoinbaseTX(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		// TODO: Send tx to network node
	}

	fmt.Println("Success!")
}
