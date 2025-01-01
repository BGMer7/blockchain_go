package routers

import (
	"fmt"
	"log"
	"mse/internal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"encoding/hex"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Server structure
type Server struct {
	nodeID string
	bch    *internal.BlockchainHandler
	wh     *internal.WalletHandler
}

func NewServer() *Server {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID env. var is not set!")
	}
	return &Server{
		wh:     internal.NewWalletHandler(),
		bch:    internal.NewBlockchainHandler(),
		nodeID: nodeID,
	}
}

func (s *Server) SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS
	r.Use(cors.Default())

	// API 路由组
	api := r.Group("/api")
	{
		// step1: create a wallet
		api.POST("/wallet/:nodeId", s.createWallet)

		// step2:create a blockchain
		api.POST("/blockchain/:address/:nodeId", s.createBlockchain)

		// step3: list the blockchain
		api.GET("/chain/:nodeId", s.printChain)

		// step4: query the balance
		api.GET("/balance/:address/:nodeId", s.getBalance)

		// step5: list the addresses
		api.GET("/addresses/:nodeId", s.listAddresses)

		// step6: transaction
		api.POST("/send/:nodeId", s.send)

		// step7: query the latest transaction
		api.GET("/latestTx/:nodeId", s.getLastTransaction)

		// step8:
		api.GET("/txCount/:nodeId", s.reindexUTXO)
	}

	// 显式处理静态文件
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.GET("/js/app.js", func(c *gin.Context) {
		c.File("./static/js/app.js")
	})

	// 处理其他静态文件
	r.Static("/static", "./static")

	return r
}

func (s *Server) getBalance(c *gin.Context) {
	address := c.Param("address")
	nodeId := c.Param("nodeId")
	balance := s.wh.GetBalance(address, nodeId)
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
	chain := s.bch.PrintChain(nodeId)
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
	address := s.wh.CreateWallet(nodeId)
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data: gin.H{
			"success": true,
			"data":    address,
		},
	})
}

func (s *Server) createBlockchain(c *gin.Context) {
	if s.bch.DBExists() {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data: gin.H{
				"success": true,
				"data":    "blockchain already exists",
			},
		})
		return
	}

	address := c.Param("address")
	nodeId := c.Param("nodeId")
	s.bch.CreateBlockchain(address, nodeId)
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
	addresses := s.wh.ListAddresses(nodeId)

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

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		log.Panic(err)
		return
	}

	s.wh.Send(req.From, req.To, req.Amount, nodeId, req.Mine)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction sent successfully",
	})
}

func (s *Server) reindexUTXO(c *gin.Context) {
	nodeId := c.Param("nodeId")
	count := s.bch.ReindexUTXO(nodeId)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    count,
	})
}

// getLastTransaction handles the request to retrieve the most recent transaction
func (s *Server) getLastTransaction(c *gin.Context) {
	nodeID := c.Param("nodeId")
	bc := s.bch.NewBlockchain(nodeID)

	lastTx := s.bch.GetLastTransaction(bc)
	if lastTx == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "No transactions found in the blockchain",
		})
		return
	}

	// 美化交易信息
	txInfo := map[string]interface{}{
		"id":        hex.EncodeToString(lastTx.ID),
		"inputs":    len(lastTx.Vin),
		"outputs":   len(lastTx.Vout),
		"timestamp": time.Now().Unix(), // 可以考虑在交易结构中添加时间戳
		"details":   lastTx.ToJSON(),
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction": txInfo,
		},
	})
}
