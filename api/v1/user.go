package v1

import (
	"banana/service"
	"banana/util"
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

func UserPost(ctx *gin.Context) {

	var userUpdate service.UserPostService
	accclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
	recclaims, _ := util.ParseToken(ctx.GetHeader("refresh_token"))
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.UserPost(ctx.Request.Context(), accclaims.ID, recclaims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}

}
