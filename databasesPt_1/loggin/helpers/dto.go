package helpers

import (
	"errors"
	"math/rand"
)

type StringsDTO struct {
	ID   int    `json:"ID"`
	Text string `json:"Text"`
}

// checks whether text is empty or not
func NewStringDTO(text string) (StringsDTO, error) {
	if text == "" {
		err := errors.New("received text is empty")
		return StringsDTO{}, err
	}
	id := rand.Intn(6969)
	return StringsDTO{
		ID:   id,
		Text: text,
	}, nil
}

type ErrorDTO struct {
	ErrorItself error `json:"Error"`
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		ErrorItself: err,
	}
}
