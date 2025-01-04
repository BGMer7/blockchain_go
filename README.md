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

### 5. 一键启动界面
```bash
./start_nodes.sh
```

## 注意事项
- 每个操作都需要指定 `NODE_ID`
- 钱包和区块链数据默认存储在 `data/` 目录
- 交易需要通过挖矿打包

## 开发和调试
- 项目使用 Go Modules 管理依赖
- 日志记录在控制台输出，方便调试

### 接口文档

其中，除命令行操作之外，该项目支持RESTful的http接口请求，由Gin实现的后端将在NodeID+30000的端口启动：

| **接口名称**      | **URL**                        | **Method** | **描述**                          | **参数**                                                     | **响应**                                                     |
| ----------------- | ------------------------------ | ---------- | --------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **创建钱包**      | `/wallet/:nodeId`              | `POST`     | 创建一个新的钱包                  | `nodeId` (路径参数): 节点 ID                                 | `200 OK`: 钱包创建成功<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **创建区块链**    | `/blockchain/:address/:nodeId` | `POST`     | 创建一个新的区块链                | `address` (路径参数): 钱包地址<br>`nodeId` (路径参数): 节点 ID | `200 OK`: 区块链创建成功<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **列出区块链**    | `/chain/:nodeId`               | `GET`      | 列出指定节点的区块链              | `nodeId` (路径参数): 节点 ID                                 | `200 OK`: 返回区块链数据<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **查询余额**      | `/balance/:address/:nodeId`    | `GET`      | 查询指定地址的余额                | `address` (路径参数): 钱包地址<br>`nodeId` (路径参数): 节点 ID | `200 OK`: 返回余额数据<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **列出地址**      | `/addresses/:nodeId`           | `GET`      | 列出指定节点的所有地址            | `nodeId` (路径参数): 节点 ID                                 | `200 OK`: 返回地址列表<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **发起交易**      | `/send/:nodeId`                | `POST`     | 发起一笔交易                      | `nodeId` (路径参数): 节点 ID<br>**请求体**:<br>`from`: 发送方地址<br>`to`: 接收方地址<br>`amount`: 交易金额 | `200 OK`: 交易成功<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **查询最新交易**  | `/latestTx/:nodeId`            | `GET`      | 查询指定节点的最新交易            | `nodeId` (路径参数): 节点 ID                                 | `200 OK`: 返回最新交易数据<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |
| **重新索引 UTXO** | `/txCount/:nodeId`             | `GET`      | 重新索引 UTXO（未花费的交易输出） | `nodeId` (路径参数): 节点 ID                                 | `200 OK`: 重新索引成功<br>`400 Bad Request`: 请求参数错误<br>`500 Internal Server Error`: 服务器内部错误 |

### 说明
- **URL** 中的 `:nodeId`、`:address` 等为路径参数，需要替换为实际值。
- **请求体** 仅适用于 `POST` 请求，需以 JSON 格式传递。
- **响应** 中的状态码表示请求结果，`200 OK` 表示成功，其他状态码表示错误。

### 示例

---

**1. 创建钱包**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 创建钱包                                                     |
| **请求示例** | `POST http://localhost:33000/wallet/3000`                    |
| **响应示例** | {<br/>    "success": true,<br/>    "data": {<br/>        "data": "1MyrDDSdyrG1ip95yfc8wazFkUiqSisJDG",<br/>        "success": true<br/>    }<br/>} |

---

**2. 创建区块链**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 创建区块链                                                   |
| **请求示例** | `POST http://localhost:33000/blockchain/1MyrDDSdyrG1ip95yfc8wazFkUiqSisJDG/3000` |
| **响应示例** | {<br/>    "success": true,<br/>    "data": {<br/>        "data": "blockchain has been successfully created.",<br/>        "success": true<br/>    }<br/>} |

---

**3. 列出区块链**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 列出区块链                                                   |
| **请求示例** | `GET http://localhost:33000/addresses/3000`                  |
| **响应示例** | {<br/>    "data": [<br/>        "1MyrDDSdyrG1ip95yfc8wazFkUiqSisJDG"<br/>    ],<br/>    "success": true<br/>} |

---

**4. 查询余额**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 查询余额                                                     |
| **请求示例** | `GET /balance/bc1q3yz7qs5dvm3l249rjqg393qfjddzev37cp84e3/node1` |
| **响应示例** | ```json { "address": "bc1q3yz7qs5dvm3l249rjqg393qfjddzev37cp84e3", "balance": 100 } ``` |

---

**5. 列出地址**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 列出地址                                                     |
| **请求示例** | `GET http://localhost:33000/balance/1MyrDDSdyrG1ip95yfc8wazFkUiqSisJDG/3000` |
| **响应示例** | {<br/>    "success": true,<br/>    "data": {<br/>        "data": 10,<br/>        "success": true<br/>    }<br/>} |

---

**6. 发起交易**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 发起交易                                                     |
| **请求示例** | `POST http://localhost:33000/send/3000`<br>**请求体**:<br>{<br/>  "from": "15yzQZyefcyQ34s2wZTgu2KbBKAXict5Ld",<br/>  "to": "1HFvegw9KNodC2pRMruewBY5RUC3Z4GZPD",<br/>  "amount": 1,<br/>  "mine": true<br/>} |
| **响应示例** |                                                              |

---

**7. 查询最新交易**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 查询最新交易                                                 |
| **请求示例** | `GET http://localhost:33000/latestTx/3000`                   |
| **响应示例** | {<br/>    "success": true,<br/>    "transaction": {<br/>        "lastTx:": {<br/>            "id": "b457ce0da8127abd295c6d03495d4c657ae727ec0506458cf2bc1865758593bc",<br/>            "vin": [<br/>                {<br/>                    "out": -1,<br/>                    "pubKey": "5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73",<br/>                    "signature": "",<br/>                    "txid": ""<br/>                }<br/>            ],<br/>            "vout": [<br/>                {<br/>                    "script": "e622e372e4199bd34ce87a07b9fb6cdf514d2d64",<br/>                    "value": 10<br/>                }<br/>            ]<br/>        }<br/>    }<br/>} |

---

**8. 重新索引 UTXO**

| **字段**     | **内容**                                                     |
| ------------ | ------------------------------------------------------------ |
| **API 名称** | 重新索引 UTXO                                                |
| **请求示例** | `GET http://localhost:33000/txCount/3000`                    |
| **响应示例** | ```json { "message": "UTXO reindexed successfully", "utxoCount": 50 } ``` |





## 许可证
MIT 许可证

## 贡献
欢迎提交 Issues 和 Pull Requests！
