package server

import (
	api "banana/api/v1"
	"banana/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "sucess")
		})
		//用户注册接口
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("/logout", api.UserLogout)
			authed.POST("/user", api.UserPost)
			authed.POST("/user/update", api.UserUpdate)
			authed.POST("/user/content", api.PostContent)
		}
	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Code": "404",
			"Msg":  "Not Found",
		})
	})
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"Code": "405",
			"Msg":  "Method Not Allowed",
		})
	})

	return r
}
