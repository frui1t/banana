package v1

import (
	"banana/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册接口
func UserRegister(ctx *gin.Context) {
	var service service.UserRegisterService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Register()
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}

func UserLogin(ctx *gin.Context) {
	var service service.UserLoginService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Login()
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}
