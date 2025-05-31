package app_error

type AppError struct {
	HttpStatus int    `json:"http_status"`
	Message    string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func New(httpStatus int, message string) AppError {
	return AppError{
		HttpStatus: httpStatus,
		Message:    message,
	}
}
