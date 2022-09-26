package server

import (
	api "banana/api/v1"
	"banana/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "sucess")
		})
		//用户注册接口
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)
		//v1.POST("/user", api.UserPost)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("/user", api.UserPost)
			authed.POST("/user/content", api.PostContent)
		}
	}

	return r
}
