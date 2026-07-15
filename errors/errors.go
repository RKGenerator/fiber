package apperrors

import "fmt"

type APIError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Status, e.Code, e.Message)
}

var (
	ErrNotFound = &APIError{
		Status:  404,
		Code:    "NOT_FOUND",
		Message: "не найдено",
	}
	ErrBadRequest = &APIError{
		Status:  400,
		Code:    "BAD_REQUEST",
		Message: "неверный запрос",
	}
	ErrInternalServer = &APIError{
		Status:  500,
		Code:    "INTERNAL_ERROR",
		Message: "внутренняя ошибка",
	}
)
