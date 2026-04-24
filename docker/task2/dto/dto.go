package dto

import (
	"errors"
	"math/rand"
)

type EmployeeDTO struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Position string `json:"position"`
}

func NewEmployeeDTO(fullName string, position string) (EmployeeDTO, error) {
	if fullName == "" || position == "" {
		err := errors.New("so either Full Name or Position are empty. You may piss off")
		return EmployeeDTO{}, err
	}
	return EmployeeDTO{
		ID:       rand.Intn(100),
		FullName: fullName,
		Position: position,
	}, nil
}

type ErrorDTO struct {
	ErrorItself error `json:"error"`
}

func NewErrorDTO(err error) ErrorDTO {
	if err == nil {
		return ErrorDTO{}
	}
	return ErrorDTO{
		ErrorItself: err,
	}
}
