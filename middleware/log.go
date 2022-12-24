package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoggerToFile() gin.HandlerFunc {
	logFilePath := viper.GetString("log.LogPath")
	logFileName := viper.GetString("log.LogName")

	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTIme := time.Now()
		latencyTime := endTIme.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()
		logger.Infof("|%3d|%13v|%15s|%s|%s|",

			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri)

	}
}
