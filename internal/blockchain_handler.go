package internal

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"mse/internal/blockchain"
	"mse/internal/wallet"
)

type BlockchainHandler struct{}

func NewBlockchainHandler() *BlockchainHandler {
	return &BlockchainHandler{}
}

func (h *BlockchainHandler) CreateBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			log.Panic("database close error:", err)
		}
	}(bc.DB)

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}

// GetLastTransaction retrieves the most recent transaction from the blockchain
func (h *BlockchainHandler) GetLastTransaction(bc *blockchain.Blockchain) *blockchain.Transaction {
	iter := bc.Iterator()

	// Traverse the blockchain starting from the latest block
	for {
		block := iter.Next()

		// Check if there are transactions in this block
		if len(block.Transactions) > 0 {
			// Return the last transaction of the block
			return block.Transactions[len(block.Transactions)-1]
		}

		// If it's the genesis block, stop searching
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return nil // No transaction found
}

func (h *BlockchainHandler) PrintChain(nodeID string) json.RawMessage {
	bc := blockchain.NewBlockchain(nodeID)
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			log.Panic("database close error:", err)
		}
	}(bc.DB)

	bci := bc.Iterator()

	var blocks []map[string]interface{}

	for {
		block := bci.Next()

		blockInfo := map[string]interface{}{
			"hash":          fmt.Sprintf("%x", block.Hash),
			"height":        block.Height,
			"prevBlockHash": fmt.Sprintf("%x", block.PrevBlockHash),
			"powValid":      blockchain.NewProofOfWork(block).Validate(),
			"transactions":  []map[string]interface{}{},
		}

		for _, tx := range block.Transactions {
			txInfo := map[string]interface{}{
				"id":      tx.ID,
				"inputs":  []map[string]interface{}{},
				"outputs": []map[string]interface{}{},
			}

			for _, input := range tx.Vin {
				inputInfo := map[string]interface{}{
					"txid":      fmt.Sprintf("%x", input.Txid),
					"out":       input.Vout,
					"signature": fmt.Sprintf("%x", input.Signature),
					"pubKey":    fmt.Sprintf("%x", input.PubKey),
				}
				txInfo["inputs"] = append(txInfo["inputs"].([]map[string]interface{}), inputInfo)
			}

			for _, output := range tx.Vout {
				outputInfo := map[string]interface{}{
					"value":  output.Value,
					"script": fmt.Sprintf("%x", output.PubKeyHash),
				}
				txInfo["outputs"] = append(txInfo["outputs"].([]map[string]interface{}), outputInfo)
			}

			blockInfo["transactions"] = append(blockInfo["transactions"].([]map[string]interface{}), txInfo)
		}

		blocks = append(blocks, blockInfo)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	// 将 blocks 转换为 JSON
	jsonData, err := json.Marshal(blocks)
	if err != nil {
		log.Panic("JSON marshaling error:", err)
	}

	// 使用 json.RawMessage 避免转义
	rawMessage := json.RawMessage(jsonData)
	finalOutput := map[string]interface{}{
		"data": rawMessage,
	}

	// 将最终的输出转换为 JSON
	finalJSON, err := json.Marshal(finalOutput)
	if err != nil {
		log.Panic("JSON marshaling error:", err)
	}

	return finalJSON
}

func (h *BlockchainHandler) ReindexUTXO(nodeID string) int {
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
	return count
}

func (h *BlockchainHandler) NewBlockchain(nodeID string) *blockchain.Blockchain {
	return blockchain.NewBlockchain(nodeID)
}
