package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

// InitLogger inicializa o logger
func InitLogger(output string, level string) {
	logConfig := zap.Config{
		OutputPaths: []string{output},
		Level:       zap.NewAtomicLevelAt(getLevelLogs(level)),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	log = logger
}

// Info registra uma mensagem de informação
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

// Error registra uma mensagem de erro
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(message, tags...)
	log.Sync()
}

func getLevelLogs(level string) zapcore.Level {
	switch level {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func SetupLogger() error {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		return err
	}
	log = logger
	return nil
}

func GetLogger() *zap.Logger {
	return log
}
