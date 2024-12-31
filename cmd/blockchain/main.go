package main

import (
	"log"
	"os"
)

func main() {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID env. var is not set!")
	}
	
	cli := CLI{nodeID: nodeID}
	cli.Run()
}
