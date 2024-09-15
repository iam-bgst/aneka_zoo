package repository

import (
	"context"
	"time"
	
	"gorm.io/gorm"
	"zoo/config"
	"zoo/constants"
	"zoo/domain/models"
)

type (
	IAnimalRepository interface {
		Create(ctx context.Context, data *models.Animal) error
		Update(ctx context.Context, id int, data *models.Animal) error
		Delete(ctx context.Context, id int) error
		GetByID(ctx context.Context, id int) (models.Animal, error)
		GetList(ctx context.Context, offset, limit int, orderBy, sort, search string) (data []models.Animal, count int64, err error)
		CheckExist(ctx context.Context, id int) bool
		constants.Common
	}
	animalRepository struct {
		Postgres *config.PostgresGorm
	}
)

func NewAnimalRepository(postgres *config.PostgresGorm) IAnimalRepository {
	return &animalRepository{Postgres: postgres}
}

func (repo *animalRepository) GetDB() *gorm.DB {
	return *repo.Postgres
}

func (repo *animalRepository) Create(ctx context.Context, data *models.Animal) error {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	return trx.WithContext(ctx).Model(&models.Animal{}).Omit("deleted_at").Create(data).Error
}

func (repo *animalRepository) Update(ctx context.Context, id int, data *models.Animal) error {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	return trx.WithContext(ctx).Model(&models.Animal{}).Where("id = ? AND deleted_at IS NULL", id).Omit("deleted_at").Updates(data).Error
}

func (repo *animalRepository) Delete(ctx context.Context, id int) error {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	return trx.WithContext(ctx).Model(models.Animal{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

func (repo *animalRepository) GetByID(ctx context.Context, id int) (data models.Animal, err error) {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	err = trx.WithContext(ctx).Model(&models.Animal{}).Where("id = ? AND deleted_at IS NULL", id).First(&data).Error
	return data, err
}

func (repo *animalRepository) GetList(ctx context.Context, offset, limit int, orderBy, sort, search string) (data []models.Animal, count int64, err error) {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	
	db := trx.WithContext(ctx).Model(&models.Animal{}).Where("deleted_at IS NULL")
	
	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}
	if orderBy != "" {
		db = db.Order(orderBy + " " + sort)
	}
	err = db.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset(offset).
		Limit(limit).
		Order(orderBy + " " + sort).
		Find(&data).Error
	return data, count, err
}

func (repo *animalRepository) CheckExist(ctx context.Context, id int) bool {
	trx, ok := ctx.Value(constants.KeyTransaction).(*gorm.DB)
	if !ok {
		trx = *repo.Postgres
	}
	var count int64
	trx.WithContext(ctx).Model(&models.Animal{}).Where("id = ? AND deleted_at IS NULL", id).Count(&count)
	return count > 0
}
