package v1

import (
	"banana/service"
	"banana/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostContent(ctx *gin.Context) {
	var service service.ContentPostService
	recclaims, _ := util.ParseToken(ctx.GetHeader("refresh_token"))
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}