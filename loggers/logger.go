package loggers

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
}

type Zap struct {
	ZapLogger *zap.Logger
}

func (z Zap) Info(msg string) {
	z.ZapLogger.Info(msg)
}

func (z Zap) Debug(msg string) {
	z.ZapLogger.Debug(msg)
}

func (z *Zap) SetZapLogger() {
	rawJSONConfig := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSONConfig, &cfg); err != nil {
		panic(err)
	}

	z.ZapLogger = zap.Must(cfg.Build())
}
