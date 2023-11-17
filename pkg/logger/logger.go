package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger zap.Logger

func init() {
	log.Printf("-->>> pkg Logger init()")
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info
	// level := zapcore.InfoLevel
	// level := zapcore.WarnLevel
	level := zap.DebugLevel
	logLevel := zap.NewAtomicLevelAt(level) // PROD set to `warn`
	config.Level = logLevel
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// can standard output, like you want to output log to some file use `>`
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	Logger = logger.Sugar()
}
