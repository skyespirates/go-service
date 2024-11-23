package utils

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	Logger = logger.Sugar()
	defer logger.Sync()
}
