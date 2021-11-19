package initialize

import (
	"mxshop_api/user-web/middlewares"
	"mxshop_api/user-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// 跨域
	Router.Use(middlewares.Cors())

	// Api分组
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return Router
}
