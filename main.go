package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = CollectRoute(r)
	r.Run("localhost:8080")
	//fmt.Printf("a")
}
