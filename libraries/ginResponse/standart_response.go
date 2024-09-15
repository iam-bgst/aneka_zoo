package ginResponse

import (
	"github.com/gin-gonic/gin"
	requestId "zoo/middleware"
)

type (
	ResponsePresenter struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		TraceId string      `json:"trace_id"`
		Data    interface{} `json:"data,omitempty"`
		Meta    interface{} `json:"meta,omitempty"`
	}
	
	MetaResponsePresenter struct {
		CurrentPage int `json:"current_page"`
		LastPage    int `json:"last_page"`
		Total       int `json:"total"`
		PerPage     int `json:"per_page"`
	}
	
	OdooResponsePresenter struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
	}
	
	IStandardResponse interface {
		SendResponseWithoutMeta(ctx *gin.Context, message string, data interface{}, httpStatus int)                // Send response without meta
		SendResponseWithMeta(ctx *gin.Context, data interface{}, message string, meta interface{}, httpStatus int) // Send response with meta
	}
)

func SendResponseWithoutMeta(ctx *gin.Context, message string, data interface{}, httpStatus int) {
	requestId := requestId.Get(ctx)
	// Success is indicated with 2xx status codes
	detectSuccess := httpStatus >= 200 && httpStatus < 300
	
	// Response standart data
	ctx.AbortWithStatusJSON(httpStatus, ResponsePresenter{
		Success: detectSuccess,
		Message: message,
		TraceId: requestId,
		Data:    data,
	})
	return
}

func SendResponseWithMeta(ctx *gin.Context, data interface{}, message string, meta interface{}, httpStatus int) {
	reqId := requestId.Get(ctx)
	// Success is indicated with 2xx status codes:
	detectSuccess := httpStatus >= 200 && httpStatus < 300
	
	// Response standart data
	ctx.AbortWithStatusJSON(httpStatus, ResponsePresenter{
		Success: detectSuccess,
		Message: message,
		TraceId: reqId,
		Data:    data,
		Meta:    meta,
	})
	return
}
