package errors

type RestError struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}


func BadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status: 400,
		Error: "bad_request",
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status: 404,
		Error: "not_found",
	}
}