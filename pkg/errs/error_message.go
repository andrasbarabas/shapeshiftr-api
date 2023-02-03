package errs

var ErrorMessages = map[int]string{
	400001: "The server cannot process the request.",
	404001: "The requested resource was not found.",
	422001: "Unprocessable entity error.",
	422002: "Field is required.",
	500001: "An internal server error occurred.",
}

func GetErrorMessage(errorCode int) string {
	message, ok := ErrorMessages[errorCode]

	if ok {
		return message
	}

	return ErrorMessages[InternalServerError]
}
