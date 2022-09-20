package server

import (
	api "banana/api/v1"

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
	}

	return r
}
