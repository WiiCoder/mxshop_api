package router

import (
	"mxshop_api/user-web/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", api.GetUserList)
		UserRouter.POST("login  ", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
		UserRouter.GET("detail", api.GetUserDetail)
		UserRouter.PATCH("update", api.UpdateUser)
	}
}
