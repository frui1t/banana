package server

import (
	api "banana/api/v1"
	"banana/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), middleware.LoggerToFile())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.Static("/images", "./images")

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "200",
				"msg":    "成功",
			})
		})
		v1.POST("/file", func(ctx *gin.Context) {
			file, _ := ctx.FormFile("file")
			dst := "images/" + file.Filename
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(200, "上传成功")

		})
		//用户注册接口
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)
		v1.POST("/content", api.GetContent)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("/logout", api.UserLogout)
			authed.POST("/user", api.UserPost)
			//更新用户信息
			authed.POST("/user/update", api.UserUpdate)
			//新建动态
			authed.POST("/post/content", api.PostContentCreate)
			// //删除动态
			// authed.DELETE("/post/content", api.PostContentDelete)
			// //发布动态评论
			// authed.POST("/post/comment", api.PostCommentCreate)
			// //删除动态评论
			// authed.POST("/post/comment", api.PostCommentDelete)
		}
	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "404",
			"msg":    "Not Found",
		})
	})
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"status": "405",
			"msg":    "Method Not Allowed",
		})
	})

	return r
}
