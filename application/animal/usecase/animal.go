package usecase

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"
	"zoo/domain/models"
	"zoo/domain/transaction"
	"zoo/libraries/ginResponse"
	
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	requestPayload "zoo/application/animal/delivery/v1http/request"
	"zoo/application/animal/delivery/v1http/response"
	animalRepository "zoo/application/animal/repository"
	"zoo/libraries/logger"
)

type (
	IAnimalUsecase interface {
		GetByID(ctx context.Context, requestId string, id string) (payloadResponse response.AnimalReadByIdResponse, code int, err error)
		Create(ctx context.Context, requestId string, payload requestPayload.AnimalCreateRequest) (code int, err error)
		Update(ctx context.Context, requestId, id string, payload requestPayload.AnimalUpdateRequest) (code int, err error)
		Delete(ctx context.Context, requestId, id string) (code int, err error)
		GetList(ctx context.Context, requestId string, payload requestPayload.ListAnimalRequest) (code int, payloadResponse []response.AnimalListResponse, meta ginResponse.MetaResponsePresenter, err error)
	}
	animalUsecase struct {
		Logger           logger.ILogger
		AnimalRepository animalRepository.IAnimalRepository
	}
)

func NewAnimalUsecase(logger logger.ILogger, animalRepository animalRepository.IAnimalRepository) IAnimalUsecase {
	return &animalUsecase{
		Logger:           logger,
		AnimalRepository: animalRepository,
	}
}

func (uc *animalUsecase) GetByID(ctx context.Context, requestId string, id string) (payloadResponse response.AnimalReadByIdResponse, code int, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "strconv.Atoi", "Failed to convert id to int")
		return payloadResponse, http.StatusInternalServerError, err
	}
	
	dataAnimal, err := uc.AnimalRepository.GetByID(ctx, idInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.Logger.ErrorWithContext(ctx, requestId, err, "GetByID", "error1", "Animal not found")
			return payloadResponse, http.StatusNotFound, err
		}
		uc.Logger.ErrorWithContext(ctx, requestId, err, "GetByID", "error1", "Failed to get animal by id")
		return payloadResponse, http.StatusForbidden, err
	}
	
	err = copier.Copy(&payloadResponse, &dataAnimal)
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "copier.Copy", "Failed to copy data")
		return payloadResponse, http.StatusConflict, err
	}
	
	return payloadResponse, http.StatusOK, nil
}

func (uc *animalUsecase) Create(ctx context.Context, requestId string, payload requestPayload.AnimalCreateRequest) (code int, err error) {
	trx, trxCtx, err := transaction.Begin(ctx, uc.AnimalRepository.GetDB())
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "transaction.Begin: error")
		return http.StatusInternalServerError, err
	}
	
	modelAnimal := &models.Animal{
		Name:      payload.Name,
		Class:     payload.Class,
		Legs:      payload.Legs,
		CreatedAt: time.Now(),
	}
	
	err = uc.AnimalRepository.Create(trxCtx, modelAnimal)
	if err != nil {
		trx.Rollback()
		uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.Create: error")
		return http.StatusForbidden, err
	}
	
	err = trx.Commit()
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "trx.Commit: error")
		return http.StatusInternalServerError, err
	}
	
	return http.StatusOK, nil
}

func (uc *animalUsecase) Update(ctx context.Context, requestId, id string, payload requestPayload.AnimalUpdateRequest) (code int, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "strconv.Atoi", "Failed to convert id to int")
		return http.StatusInternalServerError, err
	}
	
	if found := uc.AnimalRepository.CheckExist(ctx, idInt); !found {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.CheckExist", "error1", "Animal not found")
		return http.StatusNotFound, errors.New("Animal not found")
	}
	
	trx, trxCtx, err := transaction.Begin(ctx, uc.AnimalRepository.GetDB())
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "transaction.Begin: error")
		return http.StatusInternalServerError, err
	}
	
	modelAnimal := &models.Animal{
		Name:      payload.Name,
		Class:     payload.Class,
		Legs:      payload.Legs,
		CreatedAt: time.Now(),
	}
	
	err = uc.AnimalRepository.Update(trxCtx, idInt, modelAnimal)
	if err != nil {
		trx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.Update", "error1", "Animal not found")
			return http.StatusNotFound, err
		}
		uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.Update: error")
		return http.StatusForbidden, err
	}
	
	err = trx.Commit()
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "trx.Commit: error")
		return http.StatusInternalServerError, err
	}
	
	return http.StatusOK, nil
}

func (uc *animalUsecase) Delete(ctx context.Context, requestId, id string) (code int, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "strconv.Atoi", "Failed to convert id to int")
		return http.StatusInternalServerError, err
	}
	
	if found := uc.AnimalRepository.CheckExist(ctx, idInt); !found {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.CheckExist", "error1", "Animal not found")
		return http.StatusNotFound, errors.New("Animal not found")
	}
	
	trx, trxCtx, err := transaction.Begin(ctx, uc.AnimalRepository.GetDB())
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "transaction.Begin: error")
		return http.StatusInternalServerError, err
	}
	
	err = uc.AnimalRepository.Delete(trxCtx, idInt)
	if err != nil {
		trx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.Delete", "error1", "Animal not found")
			return http.StatusNotFound, err
		}
		uc.Logger.ErrorWithContext(ctx, requestId, err, "uc.AnimalRepository.Delete: error")
		return http.StatusForbidden, err
	}
	
	err = trx.Commit()
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "trx.Commit: error")
		return http.StatusInternalServerError, err
	}
	
	return http.StatusOK, nil
}

func (uc *animalUsecase) GetList(ctx context.Context, requestId string, payload requestPayload.ListAnimalRequest) (code int, payloadResponse []response.AnimalListResponse, meta ginResponse.MetaResponsePresenter, err error) {
	offset, limit, page, order, sort := ginResponse.SetPaginationParameter(payload.Page, payload.PerPage, payload.OrderBy, payload.Sort)
	
	data, count, err := uc.AnimalRepository.GetList(ctx, offset, limit, order, sort, payload.Search)
	if err != nil {
		uc.Logger.ErrorWithContext(ctx, requestId, err, "u.PajakUsecase.GetList: error")
		return
	}
	
	copier.Copy(&payloadResponse, &data)
	meta = ginResponse.SetPaginationResponse(page, limit, int(count))
	return http.StatusOK, payloadResponse, meta, nil
}
