# Go 区块链项目

## 项目简介
这是一个用 Go 语言实现的简单区块链项目，提供基本的区块链功能，包括钱包创建、交易、挖矿等。

## 功能特性
- 创建和管理钱包
- 创建区块链
- 发送交易
- 简单的工作量证明挖矿机制

## 环境准备
- Go 1.16 或更高版本
- 终端/命令行工具

## 安装依赖
```bash
go mod tidy
```

## 使用说明

### 1. 创建钱包
```bash
# 创建新钱包，NODE_ID 可以是任意数字
NODE_ID=3000 go run ./cmd/blockchain/ createwallet
```

### 2. 创建区块链
```bash
# 使用新创建的钱包地址作为创世区块的奖励地址
NODE_ID=3000 go run ./cmd/blockchain/ createblockchain -address <你的钱包地址>
```

### 3. 发送交易
```bash
# 从一个钱包发送币到另一个钱包
NODE_ID=3000 go run ./cmd/blockchain/ send -from <发送方地址> -to <接收方地址> -amount <金额>

# 使用 -mine 标志可以立即挖矿打包交易
NODE_ID=3000 go run ./cmd/blockchain/ send -from <发送方地址> -to <接收方地址> -amount <金额> -mine
```

### 4. 查看余额
```bash
NODE_ID=3000 go run ./cmd/blockchain/ getbalance -address <钱包地址>
```

## 注意事项
- 每个操作都需要指定 `NODE_ID`
- 钱包和区块链数据默认存储在 `data/` 目录
- 交易需要通过挖矿打包

## 开发和调试
- 项目使用 Go Modules 管理依赖
- 日志记录在控制台输出，方便调试

## 许可证
MIT 许可证

## 贡献
欢迎提交 Issues 和 Pull Requests！
