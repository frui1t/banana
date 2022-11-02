package v1

import (
	"banana/service/userservice"
	"banana/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册接口
func UserRegister(ctx *gin.Context) {
	var service userservice.UserRegisterService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Register()
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}

// 用户登录接口
func UserLogin(ctx *gin.Context) {
	var service userservice.UserLoginService
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Login()
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}

// 用户退出接口
func UserLogout(ctx *gin.Context) {
	var service userservice.UserLogoutService
	recclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.Logout(ctx.Request.Context(), recclaims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}

func UserPost(ctx *gin.Context) {

	var userUpdate userservice.UserPostService
	/*
		accclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
		recclaims, _ := util.ParseToken(ctx.GetHeader("refresh_token"))
	*/
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.UserPost()
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
	/*
		if err := ctx.ShouldBind(&userUpdate); err == nil {
			res := userUpdate.UserPost(ctx.Request.Context(), accclaims.ID, recclaims.ID)
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.JSON(http.StatusBadRequest, err)
		}
	*/

}

func UserUpdate(ctx *gin.Context) {
	
}
