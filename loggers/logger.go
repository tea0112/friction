package loggers

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(string, ...zap.Field)
	Error(string, ...zap.Field)
	Debug(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	WARN(string, ...zap.Field)
}

type Zap struct {
	ZapLogger *zap.Logger
}

func (z *Zap) WARN(msg string, fields ...zap.Field) {
	z.ZapLogger.Warn(msg, fields...)
}

func (z *Zap) Error(msg string, fields ...zap.Field) {
	z.ZapLogger.Error(msg, fields...)
}

func (z *Zap) Fatal(msg string, fields ...zap.Field) {
	z.ZapLogger.Fatal(msg, fields...)
}

func (z *Zap) Info(msg string, fields ...zap.Field) {
	z.ZapLogger.Info(msg, fields...)
}

func (z *Zap) Debug(msg string, fields ...zap.Field) {
	z.ZapLogger.Debug(msg, fields...)
}

func (z *Zap) SetZapLogger() {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
			"/tmp/logs",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	configBuild, err := config.Build()
	if err != nil {
		panic(err)
	}

	z.ZapLogger = zap.Must(configBuild, err)
}
