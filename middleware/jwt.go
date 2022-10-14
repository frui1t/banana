package middleware

import (
	"banana/model"
	"banana/util"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		access_token := ctx.GetHeader("access_token")
		logrus.Info(access_token)
		var code int
		code = 200
		if access_token == "" {
			code = 404
		}
		acc, err := util.ParseToken(access_token)
		if err != nil {
			code = 404
		}
		id := acc.ID
		_, err = model.RDB.Get(context.Background(), strconv.FormatInt(int64(id), 10)).Result()
		if err != nil {
			code = 303
		}
		if code == 303 {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    "redis err",
			})
			ctx.Abort()
			return
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
