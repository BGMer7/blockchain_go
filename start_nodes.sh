#!/bin/bash

# Function to start a node
start_node() {
    local node_id=$1
    NODE_ID=$node_id go run ./cmd/blockchain startnode &
}

# 清空 data 文件夹
rm -rf ./data
mkdir -p ./data

# Kill any existing blockchain nodes
pkill -f "cmd/blockchain startnode"

# Wait a moment to ensure ports are freed
sleep 1

# Start 3 nodes with different configurations
start_node 3000
start_node 3001
start_node 3002

# Wait for all background processes
wait

echo "All nodes started successfully!"
