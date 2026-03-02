package dto

type DTObject struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewDto(message string, err error) DTObject {
	return DTObject{
		Message: message,
		Error:   err,
	}
}
