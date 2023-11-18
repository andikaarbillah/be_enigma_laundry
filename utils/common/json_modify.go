package common

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"enigmacamp.com/be-enigma-laundry/utils/model_util"
)

func SendCreateResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusCreated, model_util.SingleResponse{
		Status: model_util.Status{
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, model_util.SingleResponse{
		Status: model_util.Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.JSON(code, model_util.SingleResponse{
		Status: model_util.Status{
			Code:        code,
			Description: description,
		},
	})
}

func SendPagedResponse(ctx *gin.Context, description string, data []any, paging any) {
	ctx.JSON(http.StatusOK, model_util.PagedResponse{
		Status: model_util.Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}
