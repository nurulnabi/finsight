package errors

type AppError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (e AppError) NotFoundError(msg string) *AppError {
	e.Code = 404
	e.ErrorCode = "NOT_FOUND"
	e.Message = msg
	return &e
}

func (e AppError) DbConnectionError(msg string) *AppError {
	e.Code = 500
	e.ErrorCode = "DB_CONNECTION_FAILED"
	e.Message = msg
	return &e
}
