package errs

import (
	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, e []Error) {
	errorResponse := ErrorResponse{}

	for _, err := range e {
		errorResponse.Errors = append(errorResponse.Errors, err)
		errorResponse.Status = err.status
	}

	ctx.AbortWithStatusJSON(errorResponse.Status, e)
}
