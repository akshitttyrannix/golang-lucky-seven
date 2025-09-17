package customerror

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errJSON(ctx *gin.Context, err error, status int) {
	ctx.JSON(status, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func errAbortWithStatusJSON(ctx *gin.Context, err error, status int) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"message": err.Error(),
		"status":  0,
	})
}

func BadRequest(ctx *gin.Context, err error) {
	errJSON(ctx, err, http.StatusBadRequest)
}

func NotFound(ctx *gin.Context, err error) {
	errJSON(ctx, err, http.StatusNotFound)
}

func SomethingWentWrong(ctx *gin.Context, err error) {
	errJSON(ctx, err, http.StatusInternalServerError)
}

func Unauthorized(ctx *gin.Context, err error) {
	errJSON(ctx, err, http.StatusUnauthorized)
}

func UnprocessableEntity(ctx *gin.Context, err error) {
	errJSON(ctx, err, http.StatusUnprocessableEntity)
}

func Abort(ctx *gin.Context, err error) {
	errAbortWithStatusJSON(ctx, err, http.StatusUnauthorized)
}
