package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger(level string) (*zap.Logger, error) {
	atom, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom))
	return logger, nil
}
