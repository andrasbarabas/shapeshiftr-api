package model

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"gorm.io/gorm"
)

type FoodModel interface {
	GetCreatedAt() time.Time
	GetDeletedAt() sql.NullTime
	GetID() uint
	GetRefID() RefID
	GetSymbol() string
	GetUpdatedAt() time.Time
}

type Food struct {
	gorm.Model

	ID     uint  `gorm:"primary_key;auto_increment"`
	RefID  RefID `gorm:"type:uuid;default:gen_random_uuid();index:,unique"`
	Symbol string
}

func (t *Food) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Food) GetDeletedAt() sql.NullTime {
	return sql.NullTime(t.DeletedAt)
}

func (t *Food) GetID() uint {
	return t.ID
}

func (t *Food) GetRefID() RefID {
	return t.RefID
}

func (t *Food) GetSymbol() string {
	return t.Symbol
}

func (t *Food) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *Food) TableName() string {
	return "food"
}

type FoodRequest struct {
	Symbol string `json:"symbol"`
}

func (t *FoodRequest) Validate() []errs.Error {
	var errors []errs.Error

	if t.Symbol == "" {
		errors = append(errors, *errs.CreateError(
			"Field is required.",
			http.StatusUnprocessableEntity,
			map[string]string{
				"field": "symbol",
			},
		))
	}

	return errors
}

type FoodResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	RefID     RefID     `json:"refId"`
	Symbol    string    `json:"symbol"`
}
