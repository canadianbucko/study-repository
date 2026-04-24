package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"yayca/dto"
)

func SendError(w http.ResponseWriter, status int, err error) {
	if err == nil {
		err = errors.New("internal server error")
	}

	errorDto := dto.NewErrorDTO(err)
	marshalledErrorDTO, _ := json.Marshal(errorDto) // ошибку игнорим тк очень редко маршал

	w.WriteHeader(status)
	w.Write(marshalledErrorDTO)
}
