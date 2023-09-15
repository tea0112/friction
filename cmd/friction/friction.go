package main

import (
	"database/sql"
	"errors"
	"friction/databases"
	"friction/loggers"
	"log"
	"reflect"
	"time"

	_ "github.com/lib/pq"
)

func isZap(logger *loggers.Logger) bool {
	if logger == nil {
		return false
	}

	loggerVal := *logger

	return reflect.TypeOf(loggerVal.(loggers.Zap)).Name() == "Zap"
}

func syncLogger(logger *loggers.Logger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}

	if isZap(logger) == false {
		return nil
	}

	zap := (*logger).(loggers.Zap)
	zap.ZapLogger.Sync()

	return nil
}

func main() {
	var logger loggers.Logger

	zap := loggers.Zap{}
	zap.SetZapLogger()
	logger = zap
	defer syncLogger(&logger)

	logger.Info(time.Now().GoString())
	logger.Debug(time.Now().GoString())

	var databaseConfig databases.Database
	db, err := sql.Open("postgres", databaseConfig.DBInfo())
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

}
