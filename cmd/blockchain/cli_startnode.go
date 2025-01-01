package main

import (
	"fmt"
	"log"
	"mse/routers"
	"os"
	"path/filepath"
)

func (cli *CLI) startNode(minerAddress string, nodeID string) {
	// 删除特定节点ID对应的数据库和数据文件
	dataDir := "./data"
	datFile := filepath.Join(dataDir, fmt.Sprintf("wallet_%s.dat", nodeID))

	// 删除数据文件
	if _, err := os.Stat(datFile); !os.IsNotExist(err) {
		err := os.Remove(datFile)
		if err != nil {
			log.Printf("Error removing data file: %v", err)
		} else {
			fmt.Printf("Removed data file for node %s\n", nodeID)
		}
	}

	server := routers.NewServer()
	r := server.SetupRouter()
	err := r.Run(":3" + nodeID)
	if err != nil {
		log.Panic(err)
		return
	}
}
