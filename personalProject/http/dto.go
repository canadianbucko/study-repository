package http

type errDTO struct {
	Message error `json:"message"`
}

//захуярить в dto это?

func NewErrDTO(err error) errDTO {
	return errDTO{
		Message: err,
	}

}

type responseDTO struct {
	Message string `json:"message"`
}

func NewResponseDTO(message string) responseDTO {
	return responseDTO{
		Message: message,
	}
}
