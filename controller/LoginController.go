package controller

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
)

var b string

func init() {
	b = RandomString()
}

func RandomString() string {
	a, _ := rand.Int(rand.Reader, big.NewInt(10))
	return fmt.Sprintf("%d", a)
}

func IsTrue(c *gin.Context) {
	log.Println("random value is", b)
	val := c.PostForm("val")
	//b:= fmt.Sprintf("%d",b)
	if val == b {
		b = RandomString()
		c.JSON(200, gin.H{
			"status": "true",
		})
		log.Println("if true,b=:", b)
		return
	}
	c.JSON(200, gin.H{
		"status": "false",
	})
}
