package main

type Handler struct{}

// NewHandler creates a new Handler instance
func NewHandler() *Handler {
	return &Handler{}
}

//func (h *Handler) getBalance(address, nodeID string) string {
//	if !ValidateAddress(address) {
//		log.Panic("ERROR: Address is not valid")
//	}
//	bc := NewBlockchain(nodeID)
//	UTXOSet := UTXOSet{bc}
//	defer bc.db.Close()
//
//	balance := 0
//	pubKeyHash := Base58Decode([]byte(address))
//	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
//	UTXOs := UTXOSet.FindUTXO(pubKeyHash)
//
//	for _, out := range UTXOs {
//		balance += out.Value
//	}
//
//	result := fmt.Sprintf("Balance of %s: %d\n", address, balance)
//	fmt.Printf(result)
//	return result
//}
//
//func (h *Handler) createWallet(nodeID string) string {
//	wallets, _ := NewWallets(nodeID)
//	address := wallets.CreateWallet()
//	wallets.SaveToFile(nodeID)
//
//	result := fmt.Sprintf("Your new address: %s\n", address)
//	fmt.Printf(result)
//	return address
//}
//
//func (h *Handler) createBlockchain(address, nodeID string) {
//	if !ValidateAddress(address) {
//		log.Panic("ERROR: Address is not valid")
//	}
//	bc := CreateBlockchain(address, nodeID)
//	defer bc.db.Close()
//
//	UTXOSet := UTXOSet{bc}
//	UTXOSet.Reindex()
//
//	fmt.Println("Done!")
//}

//// GetLastTransaction retrieves the most recent transaction from the blockchain
//func (h *Handler) GetLastTransaction(bc *Blockchain) *Transaction {
//	iter := bc.Iterator()
//
//	// Traverse the blockchain starting from the latest block
//	for {
//		block := iter.Next()
//
//		// Check if there are transactions in this block
//		if len(block.Transactions) > 0 {
//			// Return the last transaction of the block
//			return block.Transactions[len(block.Transactions)-1]
//		}
//
//		// If it's the genesis block, stop searching
//		if len(block.PrevBlockHash) == 0 {
//			break
//		}
//	}
//
//	return nil // No transaction found
//}

//func (h *Handler) listAddresses(nodeID string) []string {
//	wallets, err := NewWallets(nodeID)
//	if err != nil {
//		log.Panic(err)
//	}
//	addresses := wallets.GetAddresses()
//
//	for _, address := range addresses {
//		fmt.Println(address)
//	}
//
//	return addresses
//}
//
//func (h *Handler) printChain(nodeID string) string {
//	bc := NewBlockchain(nodeID)
//	defer bc.db.Close()
//
//	bci := bc.Iterator()
//
//	var output strings.Builder
//
//	for {
//		block := bci.Next()
//
//		blockInfo := fmt.Sprintf("============ Block %x ============\n", block.Hash)
//		fmt.Print(blockInfo)
//		output.WriteString(blockInfo)
//
//		blockInfo = fmt.Sprintf("Height: %d\n", block.Height)
//		fmt.Print(blockInfo)
//		output.WriteString(blockInfo)
//
//		blockInfo = fmt.Sprintf("Prev. block: %x\n", block.PrevBlockHash)
//		fmt.Print(blockInfo)
//		output.WriteString(blockInfo)
//
//		pow := NewProofOfWork(block)
//		powInfo := fmt.Sprintf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
//		fmt.Print(powInfo)
//		output.WriteString(powInfo)
//
//		for _, tx := range block.Transactions {
//			txInfo := fmt.Sprintf("%v\n", tx)
//			fmt.Print(txInfo)
//			output.WriteString(txInfo)
//		}
//
//		output.WriteString("\n\n")
//		fmt.Print("\n\n")
//
//		if len(block.PrevBlockHash) == 0 {
//			break
//		}
//	}
//
//	return output.String()
//}

//func (h *Handler) reindexUTXO(nodeID string) int {
//	bc := NewBlockchain(nodeID)
//	UTXOSet := UTXOSet{bc}
//	UTXOSet.Reindex()
//
//	count := UTXOSet.CountTransactions()
//	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
//	return count
//}

//func (h *Handler) send(from, to string, amount int, nodeID string, mineNow bool) {
//	if !ValidateAddress(from) {
//		log.Panic("ERROR: Sender address is not valid")
//	}
//	if !ValidateAddress(to) {
//		log.Panic("ERROR: Recipient address is not valid")
//	}
//
//	bc := NewBlockchain(nodeID)
//	UTXOSet := UTXOSet{bc}
//	defer bc.db.Close()
//
//	wallets, err := NewWallets(nodeID)
//	if err != nil {
//		log.Panic(err)
//	}
//	wallet := wallets.GetWallet(from)
//
//	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)
//
//	if mineNow {
//		cbTx := NewCoinbaseTX(from, "")
//		txs := []*Transaction{cbTx, tx}
//
//		newBlock := bc.MineBlock(txs)
//		UTXOSet.Update(newBlock)
//	} else {
//		sendTx(knownNodes[0], tx)
//	}
//
//	fmt.Println("Success!")
//}
