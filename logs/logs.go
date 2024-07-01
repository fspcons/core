package logs

import (
	"fmt"

	"go.uber.org/zap"
)

func MustNewLogger(serviceName string) *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		_, _ = fmt.Printf("failed to start zap logger: %s\n", err.Error())
		panic(err.Error())
	}
	return logger.With(zap.String("service", serviceName))
}

func MustNewSugaredLogger(serviceName string) *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		_, _ = fmt.Printf("failed to start zap logger: %s\n", err.Error())
		panic(err.Error())
	}
	return logger.With(zap.String("service", serviceName)).Sugar()
}
