package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const HeaderKey = "X-Request-Id"

type (
	IRequestId interface {
		Use() gin.HandlerFunc
	}
	RequestID struct{}
)

func NewRequestID() IRequestId {
	return &RequestID{}
}

func (r *RequestID) Use() gin.HandlerFunc {
	return func(ctxGin *gin.Context) {
		if ctxGin.Request.Header.Get(HeaderKey) == "" {
			ctxGin.Request.Header.Set(HeaderKey, uuid.New().String())
		}
		if ctxGin.Writer.Header().Get(HeaderKey) == "" {
			ctxGin.Writer.Header().Set(HeaderKey, ctxGin.Request.Header.Get(HeaderKey))
		}
		ctxGin.Next()
	}
}

func Get(ctxGin *gin.Context) string {
	return ctxGin.Request.Header.Get(HeaderKey)
}
