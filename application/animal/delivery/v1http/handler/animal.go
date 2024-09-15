package handler

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"zoo/application/animal/delivery/v1http/request"
	animalUsecase "zoo/application/animal/usecase"
	"zoo/libraries/ginResponse"
	"zoo/libraries/logger"
	requestId "zoo/middleware"
)

type (
	IAnimalHandler interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		GetList(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}
	animalHandler struct {
		Logger        logger.ILogger
		AnimalUsecase animalUsecase.IAnimalUsecase
	}
)

func NewAnimalHandler(log logger.ILogger, usecase animalUsecase.IAnimalUsecase) IAnimalHandler {
	return &animalHandler{
		Logger:        log,
		AnimalUsecase: usecase,
	}
}

func (h *animalHandler) Create(ctx *gin.Context) {
	reqId := requestId.Get(ctx)
	
	var payloadRequest request.AnimalCreateRequest
	if err := ctx.Bind(&payloadRequest); err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "ctx.Bind: error")
		ginResponse.SendResponseWithoutMeta(ctx, fmt.Sprintf("failed parse body: %s", err.Error()), nil, http.StatusBadRequest)
		return
	}
	code, err := h.AnimalUsecase.Create(ctx.Request.Context(), reqId, payloadRequest)
	if err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "h.AnimalUsecase.Create: error")
		ginResponse.SendResponseWithoutMeta(ctx, "Failed create animal", nil, code)
		return
	}
	
	h.Logger.InfoWithContext(ctx.Request.Context(), reqId, "Success create animal")
	ginResponse.SendResponseWithoutMeta(ctx, "Success create animal", nil, code)
}

func (h *animalHandler) Update(ctx *gin.Context) {
	reqId := requestId.Get(ctx)
	
	id := ctx.Param("id")
	
	var payloadRequest request.AnimalUpdateRequest
	if err := ctx.Bind(&payloadRequest); err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "ctx.Bind: error")
		ginResponse.SendResponseWithoutMeta(ctx, fmt.Sprintf("failed parse body: %s", err.Error()), nil, http.StatusBadRequest)
		return
	}
	code, err := h.AnimalUsecase.Update(ctx.Request.Context(), reqId, id, payloadRequest)
	if err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "h.AnimalUsecase.Update: error")
		ginResponse.SendResponseWithoutMeta(ctx, "Failed update animal", nil, code)
		return
	}
	
	h.Logger.InfoWithContext(ctx.Request.Context(), reqId, "Success update animal")
	ginResponse.SendResponseWithoutMeta(ctx, "Success update animal", nil, code)
}

func (h *animalHandler) GetList(ctx *gin.Context) {
	reqId := requestId.Get(ctx)
	
	var payloadRequest request.ListAnimalRequest
	if err := ctx.Bind(&payloadRequest); err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "ctx.Bind: error")
		ginResponse.SendResponseWithoutMeta(ctx, fmt.Sprintf("failed parse body: %s", err.Error()), nil, http.StatusBadRequest)
		return
	}
	code, payloadResponse, meta, err := h.AnimalUsecase.GetList(ctx.Request.Context(), reqId, payloadRequest)
	if err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, "h.AnimalUsecase.GetList: error")
		ginResponse.SendResponseWithoutMeta(ctx, "Failed list animal", nil, code)
		return
	}
	ginResponse.SendResponseWithMeta(ctx, payloadResponse, "Success list animal", meta, code)
}

func (h *animalHandler) GetByID(ctx *gin.Context) {
	reqId := requestId.Get(ctx)
	
	id := ctx.Param("id")
	
	response, code, err := h.AnimalUsecase.GetByID(ctx.Request.Context(), reqId, id)
	if err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, " h.AnimalUsecase.GetByID: error")
		ginResponse.SendResponseWithoutMeta(ctx, "Failed get animal", nil, code)
		return
	}
	
	h.Logger.InfoWithContext(ctx.Request.Context(), reqId, "Success get animal")
	ginResponse.SendResponseWithoutMeta(ctx, "Success get animal", response, code)
}

func (h *animalHandler) Delete(ctx *gin.Context) {
	reqId := requestId.Get(ctx)
	
	id := ctx.Param("id")
	
	code, err := h.AnimalUsecase.Delete(ctx.Request.Context(), reqId, id)
	if err != nil {
		h.Logger.ErrorWithContext(ctx.Request.Context(), reqId, err, " h.AnimalUsecase.Delete: error")
		ginResponse.SendResponseWithoutMeta(ctx, "Failed delete animal", nil, code)
		return
	}
	
	h.Logger.InfoWithContext(ctx.Request.Context(), reqId, "Success delete animal")
	ginResponse.SendResponseWithoutMeta(ctx, "Success delete animal", nil, code)
}
