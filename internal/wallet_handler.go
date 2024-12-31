package internal

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"mse/internal/blockchain"
	"mse/internal/network"
	"mse/internal/wallet"
	"mse/pkg/utils"
)

type WalletHandler struct {
}

func NewWalletHandler() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) GetBalance(address, nodeID string) int {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}

func (h *WalletHandler) CreateWallet(nodeID string) string {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	err := wallets.SaveToFile(nodeID)
	if err != nil {
		log.Panic("Saving new wallet to file failed, ", err)
		return ""
	}

	result := fmt.Sprintf("Your new address: %s\n", address)
	fmt.Printf(result)
	return address
}

func (h *WalletHandler) ListAddresses(nodeID string) []string {
	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}

	return addresses
}

func (h *WalletHandler) Send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			log.Panic("database close error", err)
		}
	}(bc.DB)

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := blockchain.NewUTXOTransaction(wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := blockchain.NewCoinbaseTX(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		network.SendTx(network.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
