package utils

import (
	"github.com/devianwahyu/efishery/domain/model"
)

type ErrorResponse struct {
	Status  string
	Message string
}

type SuccessResponse struct {
	Status  string
	Message string
	Data    model.Customer
}

type SuccessResponseArray struct {
	Status  string
	Message string
	Data    []model.Customer
}
