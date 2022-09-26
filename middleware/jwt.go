package middleware

import (
	"banana/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		access_token := ctx.GetHeader("access_token")
		refresh_token := ctx.GetHeader("refresh_token")
		logrus.Info(access_token)
		logrus.Info(refresh_token)
		var code int
		code = 200
		if access_token == "" || refresh_token == "" {
			code = 404
		}
		_, err := util.ParseToken(access_token)
		if err != nil {
			code = 404
		}
		_, err = util.ParseToken(refresh_token)
		if err != nil {
			code = 404
		}
		if code != 200 {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    "mid err",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
