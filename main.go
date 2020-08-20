package main

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"net/http"

	"math/big"
)

func main() {
	r := gin.Default()
	r.GET("/ra", func(context *gin.Context) {
		a, _ := rand.Int(rand.Reader, big.NewInt(10))
		context.JSON(http.StatusOK, gin.H{
			"value": a,
		})
	})
	r.Run("localhost:8080")
	//fmt.Printf("a")
}
