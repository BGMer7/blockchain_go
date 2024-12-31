// 从当前页面的端口提取 NODE_ID
const NODE_ID = window.location.port.slice(1);  // 去掉端口号前的 '3'

// 基础 API 地址
const BASE_URL = '/api';

// 钱包历史记录
let walletHistory = [];

// 创建钱包
document.getElementById('createWalletBtn').addEventListener('click', async () => {
    try {
        const response = await fetch(`${BASE_URL}/wallet/${NODE_ID}`, { method: 'POST' });
        const result = await response.json();
        
        // 如果钱包创建成功，保存到历史记录
        if (result.success) {
            const walletAddress = result.data.data;
            
            // 检查是否已存在相同的钱包地址
            const addressExists = walletHistory.some(entry => entry.address === walletAddress);
            
            // 如果地址不存在，则添加
            if (!addressExists) {
                walletHistory.push({
                    address: walletAddress,
                    nodeId: NODE_ID,
                    timestamp: new Date().toISOString()
                });
            }
            
            // 显示钱包创建结果和历史记录
            const resultHtml = `
                <p>钱包创建成功: ${walletAddress}</p>
                <h6>钱包历史记录:</h6>
                <ul class="list-group">
                    ${walletHistory.map(entry => `
                        <li class="list-group-item">
                            地址: ${entry.address} 
                            (节点: ${entry.nodeId}, 
                            创建时间: ${new Date(entry.timestamp).toLocaleString()})
                        </li>
                    `).join('')}
                </ul>
            `;
            
            document.getElementById('walletResult').innerHTML = resultHtml;
        } else {
            document.getElementById('walletResult').innerHTML = '钱包创建失败';
        }
    } catch (error) {
        document.getElementById('walletResult').innerHTML = `错误: ${error.message}`;
    }
});

// 创建区块链
document.getElementById('createBlockchainBtn').addEventListener('click', async () => {
    const address = document.getElementById('blockchainAddress').value;
    try {
        const response = await fetch(`${BASE_URL}/blockchain/${address}/${NODE_ID}`, { method: 'POST' });
        const result = await response.json();
        document.getElementById('blockchainResult').innerHTML = result.success 
            ? result.data.data 
            : '区块链创建失败';
    } catch (error) {
        document.getElementById('blockchainResult').innerHTML = `错误: ${error.message}`;
    }
});

// 列出区块链
document.getElementById('listChainBtn').addEventListener('click', async () => {
    try {
        const response = await fetch(`${BASE_URL}/chain/${NODE_ID}`);
        const result = await response.json();
        
        // 调试：打印完整的返回结果
        console.log('Chain Result:', result);
        
        // 尝试处理不同的数据格式
        let chainData;
        if (result.data && result.data.data) {
            // 如果 data 是对象，直接使用
            if (typeof result.data.data === 'object') {
                chainData = result.data.data;
            } 
            // 如果 data 是 JSON 字符串，解析它
            else if (typeof result.data.data === 'string') {
                try {
                    chainData = JSON.parse(result.data.data);
                } catch (parseError) {
                    chainData = result.data.data;
                }
            }
        } else {
            // 如果没有 data 字段，使用整个结果
            chainData = result;
        }
        
        // 格式化并显示数据
        const formattedChain = JSON.stringify(chainData, null, 2);
        
        document.getElementById('blockchainResult').innerHTML = `
            <pre class="bg-light p-3 rounded">${formattedChain}</pre>
        `;
    } catch (error) {
        console.error('Error fetching chain:', error);
        document.getElementById('blockchainResult').innerHTML = `错误: ${error.message}`;
    }
});

// 发起交易
document.getElementById('sendTxBtn').addEventListener('click', async () => {
    const from = document.getElementById('fromAddress').value;
    const to = document.getElementById('toAddress').value;
    const amount = parseInt(document.getElementById('amount').value, 10);  // 转换为整数
    
    // 验证输入
    if (isNaN(amount) || amount <= 0) {
        document.getElementById('txResult').innerHTML = '错误：请输入有效的交易金额';
        return;
    }

    try {
        const response = await fetch(`${BASE_URL}/send/${NODE_ID}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
                from, 
                to, 
                amount,  // 现在是整数 
                mine: true 
            })
        });
        const result = await response.json();
        document.getElementById('txResult').innerHTML = result.success 
            ? '交易发起成功' 
            : `交易发起失败: ${result.message || '未知错误'}`;
    } catch (error) {
        document.getElementById('txResult').innerHTML = `错误: ${error.message}`;
    }
});

// 查询余额
document.getElementById('getBalanceBtn').addEventListener('click', async () => {
    const address = document.getElementById('fromAddress').value;
    try {
        const response = await fetch(`${BASE_URL}/balance/${address}/${NODE_ID}`);
        const result = await response.json();
        document.getElementById('txResult').innerHTML = `余额: ${result.data.data}`;
    } catch (error) {
        document.getElementById('txResult').innerHTML = `错误: ${error.message}`;
    }
});

// 列出地址
document.getElementById('listAddressesBtn').addEventListener('click', async () => {
    try {
        const response = await fetch(`${BASE_URL}/addresses/${NODE_ID}`);
        const result = await response.json();
        document.getElementById('addressesResult').innerHTML = JSON.stringify(result.data, null, 2);
    } catch (error) {
        document.getElementById('addressesResult').innerHTML = `错误: ${error.message}`;
    }
});

// 最新交易
document.getElementById('latestTxBtn').addEventListener('click', async () => {
    try {
        const response = await fetch(`${BASE_URL}/latestTx/${NODE_ID}`);
        const result = await response.json();
        
        if (result.success && result.data && result.data.transaction) {
            const tx = result.data.transaction;
            const txHtml = `
                <div class="card">
                    <div class="card-header bg-primary text-white">最新交易详情</div>
                    <div class="card-body">
                        <div class="row">
                            <div class="col-md-6">
                                <strong>交易ID:</strong>
                                <code class="text-break">${tx.id}</code>
                            </div>
                            <div class="col-md-6">
                                <strong>时间:</strong>
                                ${new Date(tx.timestamp * 1000).toLocaleString()}
                            </div>
                        </div>
                        <hr>
                        <div class="row">
                            <div class="col-md-6">
                                <strong>输入数量:</strong>
                                <span class="badge bg-info">${tx.inputs}</span>
                            </div>
                            <div class="col-md-6">
                                <strong>输出数量:</strong>
                                <span class="badge bg-success">${tx.outputs}</span>
                            </div>
                        </div>
                        <hr>
                        <details>
                            <summary>交易详细信息</summary>
                            <pre class="bg-light p-3 rounded">${JSON.stringify(tx.details, null, 2)}</pre>
                        </details>
                    </div>
                </div>
            `;
            
            document.getElementById('latestTxResult').innerHTML = txHtml;
        } else {
            document.getElementById('latestTxResult').innerHTML = `
                <div class="alert alert-warning">
                    ${result.message || '未找到最新交易'}
                </div>
            `;
        }
    } catch (error) {
        document.getElementById('latestTxResult').innerHTML = `
            <div class="alert alert-danger">
                错误: ${error.message}
            </div>
        `;
    }
});
