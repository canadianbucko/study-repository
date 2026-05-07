package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func SendError(w http.ResponseWriter, err error, status int) {
	if err == nil {
		err = errors.New("no oopsie doopsie, no error provided")
	}
	errorDTO := NewErrorDTO(err)
	marshalledErrorDTO, _ := json.Marshal(errorDTO) // игнор ошибки тк не бывает
	w.WriteHeader(status)
	w.Write(marshalledErrorDTO)
}
