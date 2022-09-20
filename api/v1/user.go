package v1

import (
	"banana/service"

	"github.com/gin-gonic/gin"
)

// 用户注册接口
func UserRegister(ctx *gin.Context) {
	var service service.UserRegisterService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Register()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, "err")
	}
}
