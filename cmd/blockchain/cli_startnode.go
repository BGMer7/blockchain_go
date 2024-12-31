package main

import (
	"fmt"
	"log"
	"mse/internal/network"
	"mse/internal/wallet"
	"mse/routers"
	"os"
)

func (cli *CLI) startNode(minerAddress string) {
	fmt.Printf("Starting node %s\n", cli.nodeID)

	// 删除 data 文件夹
	dataDir := "./data"
	if _, err := os.Stat(dataDir); !os.IsNotExist(err) {
		err := os.RemoveAll(dataDir)
		if err != nil {
			log.Printf("Error removing data directory: %v", err)
		} else {
			fmt.Println("Removed existing data directory")
		}
	}

	// 重新创建 data 文件夹
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		log.Printf("Error creating data directory: %v", err)
	}

	if len(minerAddress) > 0 {
		if !wallet.ValidateAddress(minerAddress) {
			log.Panic("Wrong miner address!")
		}
		fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
	}

	// 启动 HTTP 服务器
	go func() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			log.Fatal("NODE_ID env. var is not set!")
		}
		server := routers.NewServer()
		r := server.SetupRouter()
		fmt.Printf("Starting HTTP server on port 3%s\n", nodeID)
		if err := r.Run(":3" + nodeID); err != nil {
			log.Panic(err)
		}
	}()

	// 启动网络节点服务
	network.StartServer(cli.nodeID, minerAddress)
}
