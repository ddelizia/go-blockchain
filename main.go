package main

import (
	"github.com/ddelizia/go-blockchain/module"
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
	"github.com/Sirupsen/logrus"
)

type TransactionRequest struct {
	Sender    int     `json:"sender"    binding:"required"`
	Recipient int     `json:"recipient" binding:"required"`
	Amount    float64 `json:"amount"    binding:"required"`
}

func main() {

	blockchain := module.BlockChain{}

	nodeNumber := rand.Intn(10000000)
	logrus.Info("Node created: ", nodeNumber)
	blockchain.Init(nodeNumber)

	r := gin.Default()
	r.POST("/mine", func(c *gin.Context) {
		blockchain.Mine(nodeNumber)
		c.JSON(200, gin.H{
			"status":  "mined",
		})
	})

	r.POST("/transaction", func(c *gin.Context) {
		var transaction TransactionRequest
		c.BindJSON(&transaction)
		t, err := blockchain.AddTransactionToCurrentBlock(transaction.Sender, transaction.Recipient, transaction.Amount )
		if err == nil {
			c.JSON(200, t)
		} else {
			c.String(http.StatusNotAcceptable, err.Error())
		}

	})

	r.GET("/chain", func(c *gin.Context) {
		c.JSON(200, blockchain)
	})

	r.Run()

}





