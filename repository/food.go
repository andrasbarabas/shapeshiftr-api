package repository

import (
	"errors"
	"net/http"

	"github.com/andrasbarabas/shapeshiftr-api/model"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/l"
	"gorm.io/gorm"
)

type FoodRepository interface {
	Create(t model.FoodModel) (model.FoodModel, *errs.Error)
	Delete(refID model.RefID) *errs.Error
	GetAll() ([]model.FoodModel, *errs.Error)
	GetOne(refID model.RefID) (model.FoodModel, *errs.Error)
	Update(refID model.RefID, t model.FoodModel) (model.FoodModel, *errs.Error)
}

type GORMFoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &GORMFoodRepository{
		db: db,
	}
}

func (r *GORMFoodRepository) Create(t model.FoodModel) (model.FoodModel, *errs.Error) {
	if err := r.db.Create(t).Error; err != nil {
		l.Logger.Error(err.Error())

		return nil, errs.CreateInternalServerError()
	}

	return t, nil
}

func (r *GORMFoodRepository) Delete(refID model.RefID) *errs.Error {
	var result model.Food

	if err := r.db.First(&result, "ref_id = ?", string(refID)).Delete(&result).Error; err != nil {
		l.Logger.Error(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.CreateError(
				errs.GetErrorMessage(errs.ResourceNotFoundError),
				http.StatusNotFound,
				map[string]string{
					"id": string(refID),
				},
			)
		}

		return errs.CreateInternalServerError()
	}

	return nil
}

func (r *GORMFoodRepository) GetAll() ([]model.FoodModel, *errs.Error) {
	var foods []model.Food

	if err := r.db.Find(&foods).Error; err != nil {
		l.Logger.Error(err.Error())

		return nil, errs.CreateInternalServerError()
	}

	result := make([]model.FoodModel, len(foods))

	for i, t := range foods {
		tCopy := t

		result[i] = &tCopy
	}

	return result, nil
}

func (r *GORMFoodRepository) GetOne(refID model.RefID) (model.FoodModel, *errs.Error) {
	var result model.Food

	if err := r.db.Where("ref_id = ?", string(refID)).First(&result).Error; err != nil {
		l.Logger.Error(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.CreateError(
				errs.GetErrorMessage(errs.ResourceNotFoundError),
				http.StatusNotFound,
				map[string]string{
					"id": string(refID),
				},
			)
		}

		return nil, errs.CreateInternalServerError()
	}

	return &result, nil
}

func (r *GORMFoodRepository) Update(refID model.RefID, t model.FoodModel) (model.FoodModel, *errs.Error) {
	var result model.Food

	if err := r.db.First(&result, "ref_id = ?", string(refID)).Error; err != nil {
		l.Logger.Error(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.CreateError(
				errs.GetErrorMessage(errs.ResourceNotFoundError),
				http.StatusNotFound,
				map[string]string{
					"id": string(refID),
				},
			)
		}

		return nil, errs.CreateInternalServerError()
	}

	result.Symbol = t.GetSymbol()

	if err := r.db.Omit("ref_id").Save(&result).Error; err != nil {
		l.Logger.Error(err.Error())

		return nil, errs.CreateInternalServerError()
	}

	return &result, nil
}
