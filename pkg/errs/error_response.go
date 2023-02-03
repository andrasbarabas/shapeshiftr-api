package errs

type ErrorResponse struct {
	Errors []Error `json:"errors"`
	Status int
}
