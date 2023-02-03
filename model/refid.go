package model

import (
	"net/http"

	"github.com/andrasbarabas/shapeshiftr-api/pkg/errs"
	"github.com/andrasbarabas/shapeshiftr-api/pkg/uuidvalidator"
)

type RefID string

func (id RefID) Stringer() string {
	return string(id)
}

func (id RefID) Validate() []errs.Error {
	var errors []errs.Error

	if id == "" {
		errors = append(errors, *errs.CreateError(
			"Field is required.",
			http.StatusUnprocessableEntity,
			map[string]string{"field": "id"},
		))
	}

	if !uuidvalidator.IsValidUUID(string(id)) {
		errors = append(errors, *errs.CreateError(
			"Field is not a valid UUID.",
			http.StatusUnprocessableEntity,
			map[string]string{"field": "id"},
		))
	}

	return errors
}
