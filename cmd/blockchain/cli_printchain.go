package main

import (
	"fmt"
	"strconv"

	"mse/internal/blockchain"
	"strings"
)

func (cli *CLI) printChain() {
	bc := blockchain.NewBlockchain(cli.nodeID)
	defer bc.DB.Close()

	bci := bc.Iterator()

	var output strings.Builder

	for {
		block := bci.Next()

		blockInfo := fmt.Sprintf("============ Block %x ============\n", block.Hash)
		fmt.Print(blockInfo)
		output.WriteString(blockInfo)

		blockInfo = fmt.Sprintf("Height: %d\n", block.Height)
		fmt.Print(blockInfo)
		output.WriteString(blockInfo)

		blockInfo = fmt.Sprintf("Prev. block: %x\n", block.PrevBlockHash)
		fmt.Print(blockInfo)
		output.WriteString(blockInfo)

		pow := NewProofOfWork(block)
		powInfo := fmt.Sprintf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		fmt.Print(powInfo)
		output.WriteString(powInfo)

		for _, tx := range block.Transactions {
			txInfo := fmt.Sprintf("%v\n", tx)
			fmt.Print(txInfo)
			output.WriteString(txInfo)
		}

		output.WriteString("\n\n")
		fmt.Print("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return output.String()
}
