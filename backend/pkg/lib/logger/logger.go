package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init() {
	env := os.Getenv("APP_ENV")

	var config zap.Config

	if env == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	
	if err != nil {
		panic(err)
	}

	log = logger
}

func InfoLog(msg string, data interface{}) {
	func() {
		log.Info(msg,
			zap.Any("data", data),
		)
	}()

}

func ErrorLog(msg string, data interface{}) {
	log.Error(msg,
		zap.Any("data", data),
	)
}

func LoggerSync() {
	log.Sync()
}
