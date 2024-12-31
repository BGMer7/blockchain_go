package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
)

const walletFile = "data/wallet_%s.dat"

// Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallet
}

// CurveParams 用于序列化椭圆曲线参数
type CurveParams struct {
	P       *big.Int
	N       *big.Int
	B       *big.Int
	Gx      *big.Int
	Gy      *big.Int
	BitSize int
}

// SerializableWallet 是一个可序列化的钱包结构
type SerializableWallet struct {
	PrivateKey struct {
		D      *big.Int
		Curve  CurveParams
		X, Y   *big.Int
	}
	PublicKey []byte
}

// FromEllipticCurve 从标准椭圆曲线转换为可序列化的曲线参数
func FromEllipticCurve(curve elliptic.Curve) CurveParams {
	params := curve.Params()
	return CurveParams{
		P:       params.P,
		N:       params.N,
		B:       params.B,
		Gx:      params.Gx,
		Gy:      params.Gy,
		BitSize: params.BitSize,
	}
}

// ToEllipticCurve 将可序列化的曲线参数转换回标准椭圆曲线
func (cp CurveParams) ToEllipticCurve() elliptic.Curve {
	return &elliptic.CurveParams{
		P:       cp.P,
		N:       cp.N,
		B:       cp.B,
		Gx:      cp.Gx,
		Gy:      cp.Gy,
		BitSize: cp.BitSize,
	}
}

// NewWallets creates Wallets and fills it from a file if it exists
func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	// Create data directory if it doesn't exist
	if err := os.MkdirAll("data", 0755); err != nil {
		return nil, err
	}

	// 即使 LoadFromFile 返回错误，也不要 panic
	_ = wallets.LoadFromFile(nodeID)

	return &wallets, nil
}

// CreateWallet adds a Wallet to Wallets
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	// 确保钱包映射已初始化
	if ws.Wallets == nil {
		ws.Wallets = make(map[string]*Wallet)
	}

	ws.Wallets[address] = wallet
	
	// 获取当前的 NODE_ID
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID = "3000" // 默认值
	}
	
	log.Printf("Creating wallet with address: %s, NODE_ID: %s", address, nodeID)
	
	// 保存钱包，忽略错误
	err := ws.SaveToFile(nodeID)
	if err != nil {
		log.Printf("Error saving wallet: %v", err)
	}

	return address
}

// GetAddresses returns an array of addresses stored in the wallet file
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

// GetWallet returns a Wallet by its address
func (ws *Wallets) GetWallet(address string) *Wallet {
	wallet, exists := ws.Wallets[address]
	if !exists {
		// 如果钱包不存在，返回 nil
		return nil
	}
	return wallet
}

// ToSerializable 将 Wallet 转换为可序列化的结构
func (w *Wallet) ToSerializable() SerializableWallet {
	sw := SerializableWallet{}
	sw.PrivateKey.D = w.PrivateKey.D
	sw.PrivateKey.Curve = FromEllipticCurve(w.PrivateKey.Curve)
	sw.PrivateKey.X = w.PrivateKey.X
	sw.PrivateKey.Y = w.PrivateKey.Y
	sw.PublicKey = w.PublicKey
	return sw
}

// ToWallet 将可序列化结构转换回 Wallet
func (sw SerializableWallet) ToWallet() *Wallet {
	w := &Wallet{}
	w.PrivateKey = ecdsa.PrivateKey{
		D: sw.PrivateKey.D,
		PublicKey: ecdsa.PublicKey{
			Curve: sw.PrivateKey.Curve.ToEllipticCurve(),
			X:     sw.PrivateKey.X,
			Y:     sw.PrivateKey.Y,
		},
	}
	w.PublicKey = sw.PublicKey
	return w
}

// LoadFromFile loads wallets from a file
func (ws *Wallets) LoadFromFile(nodeID string) error {
	walletFile := fmt.Sprintf(walletFile, nodeID)
	
	log.Printf("Loading wallets from file: %s", walletFile)
	
	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(walletFile), 0755); err != nil {
		log.Printf("Error creating wallet directory: %v", err)
		return err
	}

	// 如果文件不存在，初始化空的钱包映射并返回
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		log.Printf("Wallet file does not exist, initializing empty wallet set: %s", walletFile)
		ws.Wallets = make(map[string]*Wallet)
		return nil
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Printf("Error reading wallet file %s: %v", walletFile, err)
		ws.Wallets = make(map[string]*Wallet)
		return err
	}

	// 如果文件为空，初始化空的钱包映射
	if len(fileContent) == 0 {
		log.Println("Wallet file is empty, initializing empty wallet set")
		ws.Wallets = make(map[string]*Wallet)
		return nil
	}

	// 重置钱包映射
	ws.Wallets = make(map[string]*Wallet)

	// 创建一个可反序列化的结构
	var loadedData struct {
		Wallets map[string]SerializableWallet
	}

	gob.Register(CurveParams{})
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	
	err = decoder.Decode(&loadedData)
	if err != nil {
		log.Printf("Error decoding wallet file %s: %v", walletFile, err)
		return nil
	}

	// 转换回 Wallet 类型
	for addr, sw := range loadedData.Wallets {
		wallet := sw.ToWallet()
		ws.Wallets[addr] = wallet
	}

	log.Printf("Successfully loaded %d wallets from %s", len(ws.Wallets), walletFile)
	return nil
}

// SaveToFile saves wallets to a file
func (ws *Wallets) SaveToFile(nodeID string) error {
	walletFile := fmt.Sprintf(walletFile, nodeID)
	
	log.Printf("Saving %d wallets to file: %s", len(ws.Wallets), walletFile)

	// 如果没有钱包，不执行保存
	if len(ws.Wallets) == 0 {
		log.Println("No wallets to save")
		return nil
	}

	var content bytes.Buffer

	// 注册必要的类型
	gob.Register(CurveParams{})

	// 转换为可序列化的钱包
	serializableWallets := make(map[string]SerializableWallet)
	for addr, wallet := range ws.Wallets {
		serializableWallets[addr] = wallet.ToSerializable()
	}

	// 创建一个新的可序列化的 Wallets 结构
	serializableWs := struct {
		Wallets map[string]SerializableWallet
	}{
		Wallets: serializableWallets,
	}

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(serializableWs)
	if err != nil {
		log.Printf("Error encoding wallets: %v", err)
		return err
	}

	// 确保目录存在
	dir := filepath.Dir(walletFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("Error creating wallet directory: %v", err)
		return err
	}

	// 写入文件
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Printf("Error writing wallet file: %v", err)
		return err
	}

	log.Printf("Successfully saved %d wallets to %s", len(ws.Wallets), walletFile)
	return nil
}
