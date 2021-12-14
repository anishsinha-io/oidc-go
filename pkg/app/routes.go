package app

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes() {
	oauthRoutes := Router.Group("/oauth")
	{
		oauthRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		oauthRoutes.GET("/authorize", handleShowUniversalLogin)
		oauthRoutes.POST("/login", handleLogin)
		oauthRoutes.POST("/token", handleGetToken)
		oauthRoutes.POST("/introspect", handleTokenIntrospect)
		oauthRoutes.POST("/signup", handleSignup)
	}
}
