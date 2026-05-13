package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func SendError(w http.ResponseWriter, r *http.Request, err error, status int, logger *zap.Logger, startTime time.Time) {
	if err == nil {
		err = errors.New("no oopsie doopsie, no error provided")
		logger.Error("no fucking error provided jaja")
	}
	errorDTO := NewErrorDTO(err)
	marshalledErrorDTO, _ := json.Marshal(errorDTO) // игнор ошибки тк не бывает
	w.WriteHeader(status)
	w.Write(marshalledErrorDTO)

	// logger part
	logger.Error("Request failed",
		zap.Time("started_at", startTime), //snake_case!
		zap.Duration("time_since", time.Since(startTime)),
		zap.String("endpoint", r.URL.Path),
		zap.String("method", r.Method),
		zap.Int("status_code", status),
		zap.Error(err),
	)

}

//func WriteLog()

// helpers.SendError(w, err, http.StatusMethodNotAllowed)
// logger.Error("bad method",
// 	zap.Time("time", startTime),
// 	zap.String("endpoint", r.URL.Path),
// 	zap.String("method", r.Method),
// 	zap.String("statuscode", http.StatusMethodNotAllowed),
// 	zap.Time("time since", time.Since(startTime)),
// )
