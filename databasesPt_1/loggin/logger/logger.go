package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel string) (*zap.Logger, func() error, error) {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, nil, fmt.Errorf("unmarshall error log level, %w", err)
	}
	// 7 		5 				   5
	//user	 user group	     	other ppl
	// r w x		r w  x		r  w  x
	// 1 1	1		1 0 1		1  0  1
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("error creating mkdir folder, %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join("logs", fmt.Sprintf("%s.log", timestamp))

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644) // 6 r1 w1 x0     4 r1 w0 x0
	if err != nil {
		return nil, nil, fmt.Errorf("open log file error %w", err)
	}

	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15-04-05.000000")

	encoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lvl),
		zapcore.NewCore(encoder, zapcore.AddSync(logFile), lvl),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	return logger, logFile.Close, nil
}
