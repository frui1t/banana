package v1

import (
	"banana/service/userservice"
	"banana/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostContentCreate(ctx *gin.Context) {
	var service userservice.ContentPostService
	recclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}

// func PostContentDelete(ctx *gin.Context) {
// 	var service userservice.ContentPostService
// 	recclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
// 	if err := ctx.ShouldBind(&service); err == nil {
// 		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
// 		ctx.JSON(http.StatusOK, res)
// 	} else {
// 		ctx.JSON(http.StatusBadRequest, "err")
// 	}
// }

// func PostCommentCreate(ctx *gin.Context) {
// 	var service userservice.ContentPostService
// 	recclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
// 	if err := ctx.ShouldBind(&service); err == nil {
// 		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
// 		ctx.JSON(http.StatusOK, res)
// 	} else {
// 		ctx.JSON(http.StatusBadRequest, "err")
// 	}
// }

// func PostCommentDelete(ctx *gin.Context) {
// 	var service userservice.ContentPostService
// 	recclaims, _ := util.ParseToken(ctx.GetHeader("access_token"))
// 	if err := ctx.ShouldBind(&service); err == nil {
// 		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
// 		ctx.JSON(http.StatusOK, res)
// 	} else {
// 		ctx.JSON(http.StatusBadRequest, "err")
// 	}
// }
