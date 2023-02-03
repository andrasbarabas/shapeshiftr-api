package service

import (
	"github.com/andrasbarabas/shapeshiftr-api/model"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"github.com/andrasbarabas/shapeshiftr-api/repository"
)

type FoodService interface {
	Create(t model.FoodRequest) (*model.FoodResponse, *errs.Error)
	Delete(refID model.RefID) *errs.Error
	GetAll() (*[]model.FoodResponse, *errs.Error)
	GetOne(refID model.RefID) (*model.FoodResponse, *errs.Error)
	Update(refID model.RefID, t model.FoodRequest) (*model.FoodResponse, *errs.Error)
}

type GORMFoodService struct {
	repository repository.FoodRepository
}

func NewFoodService(r repository.FoodRepository) FoodService {
	return &GORMFoodService{
		repository: r,
	}
}

func (s *GORMFoodService) Create(t model.FoodRequest) (*model.FoodResponse, *errs.Error) {
	food := &model.Food{
		Symbol: t.Symbol,
	}
	result, err := s.repository.Create(food)

	if err != nil {
		return nil, err
	}

	response := &model.FoodResponse{
		CreatedAt: result.GetCreatedAt(),
		RefID:     result.GetRefID(),
		Symbol:    result.GetSymbol(),
	}

	return response, nil
}

func (s *GORMFoodService) Delete(refID model.RefID) *errs.Error {
	if err := s.repository.Delete(refID); err != nil {
		return err
	}

	return nil
}

func (s *GORMFoodService) GetAll() (*[]model.FoodResponse, *errs.Error) {
	result, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	response := make([]model.FoodResponse, len(result))

	for i, t := range result {
		response[i] = model.FoodResponse{
			CreatedAt: t.GetCreatedAt(),
			RefID:     t.GetRefID(),
			Symbol:    t.GetSymbol(),
		}
	}

	return &response, nil
}

func (s *GORMFoodService) GetOne(refID model.RefID) (*model.FoodResponse, *errs.Error) {
	result, err := s.repository.GetOne(refID)

	if err != nil {
		return nil, err
	}

	response := &model.FoodResponse{
		CreatedAt: result.GetCreatedAt(),
		RefID:     result.GetRefID(),
		Symbol:    result.GetSymbol(),
	}

	return response, nil
}

func (s *GORMFoodService) Update(refID model.RefID, t model.FoodRequest) (*model.FoodResponse, *errs.Error) {
	food := &model.Food{
		Symbol: t.Symbol,
	}
	result, err := s.repository.Update(refID, food)

	if err != nil {
		return nil, err
	}

	response := &model.FoodResponse{
		CreatedAt: result.GetCreatedAt(),
		RefID:     result.GetRefID(),
		Symbol:    result.GetSymbol(),
	}

	return response, nil
}
