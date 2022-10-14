package v1

import (
	"banana/service/userservice"
	"banana/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostContent(ctx *gin.Context) {
	var service userservice.ContentPostService
	recclaims, _ := util.ParseToken(ctx.GetHeader("Access_token"))
	if err := ctx.ShouldBind(&service); err == nil {
		res := service.ContentPost(ctx.Request.Context(), recclaims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, "err")
	}
}
