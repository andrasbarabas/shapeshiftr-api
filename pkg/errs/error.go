package errs

import (
	"net/http"
)

type Error struct {
	Message string            `json:"message"`
	Params  map[string]string `json:"params,omitempty"`
	status  int
}

func CreateInternalServerError() *Error {
	err := Error{
		Message: GetErrorMessage(InternalServerError),
		status:  http.StatusInternalServerError,
	}

	return &err
}

func CreateError(m string, s int, p map[string]string) *Error {
	err := Error{
		Message: m,
		status:  s,
	}

	if p != nil {
		err.Params = p
	}

	return &err
}

func (e *Error) Error() string {
	return e.Message
}
