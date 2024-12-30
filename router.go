package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// Request/Response structures
type SendRequest struct {
	From   string `json:"from" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
	Mine   bool   `json:"mine"`
}

type CreateBlockchainRequest struct {
	Address string `json:"address" binding:"required"`
}

type BalanceRequest struct {
	Address string `json:"address" binding:"required"`
}

type NodeStartRequest struct {
	MinerAddress string `json:"minerAddress"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Server structure
type Server struct {
	h      *Handler
	nodeID string
}

func NewServer() *Server {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID env. var is not set!")
	}
	return &Server{
		h:      NewHandler(),
		nodeID: nodeID,
	}
}

func (s *Server) setupRouter() *gin.Engine {
	r := gin.Default()

	// step1: create a wallet
	r.POST("/wallet/:nodeId", s.createWallet)

	// step2:create a blockchain
	r.POST("/blockchain/:address/:nodeId", s.createBlockchain)

	// step3: list the blockchain
	r.GET("/chain/:nodeId", s.printChain)

	// step4: query the balance
	r.GET("/balance/:address/:nodeId", s.getBalance)

	// step5: list the addresses
	r.GET("/addresses/:nodeId", s.listAddresses)

	// step6: transaction
	r.POST("/send/:nodeId", s.send)

	// step7: query the latest transaction
	r.GET("/latestTx/:nodeId", s.getLastTransaction)

	// step8:
	r.GET("/txCount/:nodeId", s.reindexUTXO)

	return r
}

func (s *Server) getBalance(c *gin.Context) {
	address := c.Param("address")
	nodeId := c.Param("nodeId")
	balance := s.h.getBalance(address, nodeId)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: gin.H{
			"success": true,
			"data":    balance,
		},
	})
}

func (s *Server) printChain(c *gin.Context) {
	nodeId := c.Param("nodeId")
	fmt.Println("nodeId: ", nodeId)
	chain := s.h.printChain(nodeId)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: gin.H{
			"success": true,
			"data":    chain,
		},
	})
}

func (s *Server) createWallet(c *gin.Context) {
	nodeId := c.Param("nodeId")
	address := s.h.createWallet(nodeId)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: gin.H{
			"success": true,
			"data":    address,
		},
	})
}

func (s *Server) createBlockchain(c *gin.Context) {
	address := c.Param("address")
	nodeId := c.Param("nodeId")
	s.h.createBlockchain(address, nodeId)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: gin.H{
			"success": true,
			"data":    "blockchain has been successfully created.",
		},
	})
}

func (s *Server) listAddresses(c *gin.Context) {
	nodeId := c.Param("nodeId")
	addresses := s.h.listAddresses(nodeId)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    addresses,
	})
}

func (s *Server) send(c *gin.Context) {
	nodeId := c.Param("nodeId")
	var req struct {
		From   string `json:"from" binding:"required"`        // 必须字段
		To     string `json:"to" binding:"required"`          // 必须字段
		Amount int    `json:"amount" binding:"required,gt=0"` // 必须字段，必须大于 0
		Mine   bool   `json:"mine,omitempty"`                 // 可选字段，默认值为 false
	}

	// 解析请求体中的 JSON 数据
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		log.Panic(err)
		return
	}

	// 处理交易逻辑，例如调用发送交易的函数
	s.h.send(req.From, req.To, req.Amount, nodeId, req.Mine)

	// 模拟成功的响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction sent successfully",
	})
}

func (s *Server) reindexUTXO(c *gin.Context) {
	nodeId := c.Param("nodeId")
	count := s.h.reindexUTXO(nodeId)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    count,
	})
}

// getLastTransaction handles the request to retrieve the most recent transaction
func (s *Server) getLastTransaction(c *gin.Context) {
	nodeID := c.Param("nodeId")
	bc := NewBlockchain(nodeID)

	lastTx := s.h.GetLastTransaction(bc)
	if lastTx == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "No transactions found in the blockchain",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"transaction": gin.H{
			"lastTx:": lastTx.String(),
		},
	})
}
