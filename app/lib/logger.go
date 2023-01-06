package lib

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	*zap.SugaredLogger
}

type GinLogger struct {
	*Logger
}

type GormLogger struct {
	*Logger
	gormlogger.Config
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger(NewEnv())
		globalLogger = &logger
	}
	return *globalLogger
}

func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func newLogger(c Env) Logger {

	config := zap.NewDevelopmentConfig()
	logOutput := c.LogOutput

	if c.Environment == "development" {
		fmt.Println("encode level")
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	if c.Environment == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	logLevel := c.LogLevel
	level := zap.PanicLevel

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.PanicLevel
	}

	config.Level.SetLevel(level)

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}
