package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

const uName, psw = "123", "123"

type Login struct {
	Username string `form:"username" binding:"username"`
	Password string `form:"username" binding:"username"`
}

type User struct {
	Username string
}

var authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
	Key:         []byte{},
	Timeout:     time.Minute * 3,
	IdentityKey: "id",
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := login.Username
		password := login.Password
		if userID == uName && password == psw {
			return &User{
				Username: userID,
			}, nil
		}
		return nil, jwt.ErrFailedAuthentication
	},
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*User); ok {
			return jwt.MapClaims{
				"id": v.Username,
			}
		}
		return jwt.MapClaims{}
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*User); ok && v.Username == uName {
			return true
		}
		return false
	},
	IdentityHandler: func(context *gin.Context) interface{} {
		claims := jwt.ExtractClaims(context)
		return &User{
			Username: claims["id"].(string),
		}
	},
})
