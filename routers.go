package main

import (
	"gin_learn/common"
	"gin_learn/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/login", common.AuthMiddleware.LoginHandler)
	r.POST("/game", common.AuthMiddleware.MiddlewareFunc(), controller.IsTrue)

	return r
}
