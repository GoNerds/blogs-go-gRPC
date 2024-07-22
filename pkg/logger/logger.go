package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(level string) {
	var zapLevel zapcore.Level

	switch level {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	default:
		zapLevel = zap.InfoLevel
	}

	config := zap.Config{
		Level: zap.NewAtomicLevelAt(zapLevel),
		Development: true,
		Encoding : "jsong",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey: "time",
			LevelKey: "level",
			NameKey: "name",
			CallerKey: "caller",
			MessageKey: "message",
			StacktraceKey: "stacktrace",
			LineEnding: zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.LowercaseLevelEncoder(),
			EncodeTime: zapcore.ISO86017TimeEncoder(),
			EncodeDuration: zapcore.StringDurationEncoder(),
			EncodeCaller: zapcore.ShortCallerEncoder(),
		},
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		log.Fatalf("Cannot initialize logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}