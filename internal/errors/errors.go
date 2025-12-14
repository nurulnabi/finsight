package errors

import (
	"log"
	"runtime/debug"
)

type AppError struct {
}

type APIError struct {
	Code      int       `json:"code"`
	ErrorCode ErrorCode `json:"error_code"`
	Message   string    `json:"message"`
}

type ErrorCode string

const (
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeInvalidInput ErrorCode = "INVALID_INPUT"
	ErrCodeInternal     ErrorCode = "INTERNAL_ERROR"
)

func (e AppError) NotFoundError(msg string) APIError {
	return APIError{
		Code:      404,
		ErrorCode: ErrCodeNotFound,
		Message:   msg,
	}
}

func (e AppError) DbConnectionError(msg string) APIError {
	return APIError{
		Code:      500,
		ErrorCode: ErrCodeInternal,
		Message:   msg,
	}
}

func (e *APIError) LogError() {
	log.Printf("err=%v\nstack=%s", e, debug.Stack())
}

func (e APIError) Error() string {
	return e.Message
}
