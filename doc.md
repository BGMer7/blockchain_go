# Go 区块链项目 - 文件和目录结构说明

## 项目目录结构

```
blockchain_go/
├── cmd/                    # 命令行入口点
│   └── blockchain/
│       ├── main.go        # 程序主入口
│       ├── cli.go         # 命令行接口定义
│       ├── cli_createblockchain.go  # 创建区块链命令
│       ├── cli_createwallet.go      # 创建钱包命令
│       ├── cli_getbalance.go        # 查询余额命令
│       ├── cli_listaddress.go       # 列出地址命令
│       ├── cli_printchain.go        # 打印区块链命令
│       ├── cli_reindexutxo.go       # 重建UTXO索引命令
│       ├── cli_send.go              # 发送交易命令
│       └── cli_startnode.go         # 启动节点命令

├── internal/               # 内部实现包
│   ├── blockchain/         # 区块链核心实现
│   │   ├── block.go        # 区块数据结构
│   │   ├── blockchain.go   # 区块链核心逻辑
│   │   ├── blockchain_iterator.go  # 区块链迭代器
│   │   ├── transaction.go  # 交易实现
│   │   ├── transaction_input.go    # 交易输入
│   │   ├── transaction_output.go   # 交易输出
│   │   ├── merkle_tree.go  # 默克尔树实现
│   │   ├── proofofwork.go  # 工作量证明算法
│   │   └── utxo_set.go     # UTXO集管理
│   
│   ├── network/            # 网络层
│   │   └── server.go       # 网络服务器实现
│   
│   └── wallet/             # 钱包管理
│       ├── wallet.go       # 单个钱包实现
│       └── wallets.go      # 多钱包管理

├── go.mod                  # Go 模块依赖管理
├── go.sum                  # 依赖校验
├── README.md               # 项目说明文档
└── doc.md                  # 详细文档

## 关键文件说明

### 命令行入口 (`cmd/blockchain/`)

#### `main.go`
- 程序启动入口
- 初始化命令行界面

#### `cli.go`
- 定义命令行接口
- 解析和路由用户命令

### 区块链核心 (`internal/blockchain/`)

#### `blockchain.go`
- 区块链主要逻辑
- 区块链创建和管理

#### `transaction.go`
- 交易数据结构
- 交易创建和验证逻辑

#### `proofofwork.go`
- 工作量证明算法实现
- 区块挖矿难度计算

### 钱包管理 (`internal/wallet/`)

#### `wallet.go`
- 单个钱包实现
- 密钥对生成
- 地址管理

#### `wallets.go`
- 多钱包管理
- 钱包持久化

### 网络层 (`internal/network/`)

#### `server.go`
- 网络通信服务器
- 节点间通信

## 关键特性

- 工作量证明共识
- UTXO 交易模型
- 公钥加密
- 多节点支持
- 钱包管理
