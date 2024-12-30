# Go 区块链项目 - 文件和目录结构说明

## 项目目录结构

```
blockchain_go/
├── cmd/                    # 命令行入口点
│   └── blockchain/
│       ├── cli.go         # 命令行接口定义
│       ├── cli_createblockchain.go  # 创建区块链命令
│       ├── cli_createwallet.go      # 创建钱包命令
│       ├── cli_getbalance.go        # 查询余额命令
│       ├── cli_listaddresses.go     # 列出地址命令
│       ├── cli_printchain.go        # 打印区块链命令
│       ├── cli_reindexutxo.go       # 重建UTXO索引命令
│       ├── cli_send.go              # 发送交易命令
│       └── main.go                  # 程序主入口

├── internal/               # 内部实现包
│   ├── blockchain/         # 区块链核心实现
│   │   ├── block.go        # 区块数据结构定义
│   │   ├── blockchain.go   # 区块链核心逻辑
│   │   ├── iterator.go     # 区块链迭代器
│   │   └── proofofwork.go  # 工作量证明算法实现
│   
│   ├── transaction/        # 交易相关实现
│   │   ├── transaction.go  # 交易数据结构和基本逻辑
│   │   └── utxo.go         # 未花费交易输出管理
│   
│   └── wallet/             # 钱包管理
│       ├── wallet.go       # 钱包数据结构和密钥生成
│       └── wallets.go      # 多钱包管理和持久化

├── pkg/                    # 公共工具包
│   └── utils/              # 通用工具函数
│       └── utils.go        # 各种辅助工具函数

└── data/                   # 数据存储目录
    ├── blockchain_*.dat    # 区块链数据文件
    └── wallet_*.dat        # 钱包数据文件
```

## 文件详细说明

### 命令行入口 (`cmd/blockchain/`)

#### `main.go`
- 程序主入口
- 初始化命令行界面并处理用户命令

#### `cli.go`
- 定义命令行接口 `CLI` 结构体
- 解析和路由不同的命令行指令
- 包含各种命令的处理方法

#### 其他 `cli_*.go` 文件
- 实现特定命令的逻辑
- 如创建钱包、发送交易、查询余额等

### 区块链核心 (`internal/blockchain/`)

#### `block.go`
- 定义 `Block` 结构体
- 包含区块的基本属性：
  - 时间戳
  - 前一个区块哈希
  - 当前区块哈希
  - 交易列表

#### `blockchain.go`
- 实现区块链的核心逻辑
- 创建创世区块
- 添加新区块
- 查找特定交易
- 处理 UTXO 集

#### `iterator.go`
- 实现区块链迭代器
- 允许按顺序遍历区块链中的区块

#### `proofofwork.go`
- 实现工作量证明（PoW）算法
- 计算区块哈希
- 验证区块是否有效

### 交易管理 (`internal/transaction/`)

#### `transaction.go`
- 定义交易结构
- 创建新交易
- 签名和验证交易
- 计算交易哈希

#### `utxo.go`
- 管理未花费交易输出（UTXO）
- 查找可用于支付的输出
- 重建 UTXO 集

### 钱包管理 (`internal/wallet/`)

#### `wallet.go`
- 生成公私钥对
- 创建比特币地址
- 签名和验证交易

#### `wallets.go`
- 管理多个钱包
- 持久化和加载钱包
- 序列化和反序列化钱包数据

### 公共工具 (`pkg/utils/`)

#### `utils.go`
- 提供通用的辅助函数
- 如数据转换、编码等实用方法

## 数据存储

### 区块链数据 (`data/blockchain_*.dat`)
- 使用 BoltDB 存储区块链数据
- 每个节点有独立的数据文件

### 钱包数据 (`data/wallet_*.dat`)
- 存储钱包信息
- 使用 gob 编码序列化

## 关键特性

- 工作量证明共识机制
- 公钥密码学
- UTXO 交易模型
- 多节点支持
- 持久化存储

## 注意事项

- 项目仅供学习和研究使用
- 非生产级区块链实现
- 缺少完整的网络和共识层

## 许可证

MIT 许可证
