package model

import (
	"strings"

	"github.com/go-playground/validator"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ValidatorResponse struct {
	Tag         string `json:"tag"`
	FailedField string `json:"field"`
	Value       any    `json:"value"`
}

type PaginationData[T any] struct {
	TotalItems int64 `json:"total_items"`
	Items      []T   `json:"items"`
}

func NewPaginationData[T any](items []T, totalItems int64) PaginationData[T] {
	return PaginationData[T]{
		Items:      items,
		TotalItems: totalItems,
	}
}

func NewSuccessResponse(data any, args ...any) *Response {
	response := &Response{
		Status: 200,
		Data:   data,
	}

	populateRes(response, args...)

	return response
}

func NewErrorResponse(err error, args ...any) *Response {
	response := &Response{
		Status: 400,
	}

	if err != nil {
		response.Message = err.Error()
	}

	populateRes(response, args...)

	return response
}

func CheckValidateErrors(errs error) *Response {
	validationErrors := []ValidatorResponse{}
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidatorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()

			validationErrors = append(validationErrors, elem)
		}
	}

	response := &Response{
		Status:  422,
		Message: "Failed validator body data please check the errors and try again",
		Data:    validationErrors,
	}

	return response
}

func populateRes(response *Response, args ...any) *Response {
	if len(args) >= 1 {
		response.Status = args[0].(int)
	}

	if len(args) == 2 {
		response.Message = strings.Join([]string{response.Message, args[1].(string)}, " ")
	}

	return response
}
