// 基础 API 地址
const BASE_URL = '/api';

// 钱包历史记录
let walletHistory = [];

// 创建钱包
document.getElementById('createWalletBtn').addEventListener('click', async () => {
    const nodeId = document.getElementById('nodeIdWallet').value;
    try {
        const response = await fetch(`${BASE_URL}/wallet/${nodeId}`, { method: 'POST' });
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
                    nodeId: nodeId,
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
    const nodeId = document.getElementById('nodeIdBlockchain').value;
    try {
        const response = await fetch(`${BASE_URL}/blockchain/${address}/${nodeId}`, { method: 'POST' });
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
    const nodeId = document.getElementById('nodeIdBlockchain').value;
    try {
        const response = await fetch(`${BASE_URL}/chain/${nodeId}`);
        const result = await response.json();
        document.getElementById('blockchainResult').innerHTML = JSON.stringify(result.data, null, 2);
    } catch (error) {
        document.getElementById('blockchainResult').innerHTML = `错误: ${error.message}`;
    }
});

// 发起交易
document.getElementById('sendTxBtn').addEventListener('click', async () => {
    const from = document.getElementById('fromAddress').value;
    const to = document.getElementById('toAddress').value;
    const amount = document.getElementById('amount').value;
    const nodeId = document.getElementById('nodeIdTx').value;
    try {
        const response = await fetch(`${BASE_URL}/send/${nodeId}`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ from, to, amount, mine: true })
        });
        const result = await response.json();
        document.getElementById('txResult').innerHTML = result.success 
            ? '交易发起成功' 
            : '交易发起失败';
    } catch (error) {
        document.getElementById('txResult').innerHTML = `错误: ${error.message}`;
    }
});

// 查询余额
document.getElementById('getBalanceBtn').addEventListener('click', async () => {
    const address = document.getElementById('fromAddress').value;
    const nodeId = document.getElementById('nodeIdTx').value;
    try {
        const response = await fetch(`${BASE_URL}/balance/${address}/${nodeId}`);
        const result = await response.json();
        document.getElementById('txResult').innerHTML = `余额: ${result.data.data}`;
    } catch (error) {
        document.getElementById('txResult').innerHTML = `错误: ${error.message}`;
    }
});

// 列出地址
document.getElementById('listAddressesBtn').addEventListener('click', async () => {
    const nodeId = document.getElementById('nodeIdAddresses').value;
    try {
        const response = await fetch(`${BASE_URL}/addresses/${nodeId}`);
        const result = await response.json();
        document.getElementById('addressesResult').innerHTML = JSON.stringify(result.data, null, 2);
    } catch (error) {
        document.getElementById('addressesResult').innerHTML = `错误: ${error.message}`;
    }
});

// 最新交易
document.getElementById('latestTxBtn').addEventListener('click', async () => {
    const nodeId = document.getElementById('nodeIdLatestTx').value;
    try {
        const response = await fetch(`${BASE_URL}/latestTx/${nodeId}`);
        const result = await response.json();
        document.getElementById('latestTxResult').innerHTML = JSON.stringify(result.data, null, 2);
    } catch (error) {
        document.getElementById('latestTxResult').innerHTML = `错误: ${error.message}`;
    }
});
