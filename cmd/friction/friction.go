package main

import (
	"database/sql"
	"errors"
	"fmt"
	"friction/databases"
	_ "friction/handlers"
	"friction/loggers"
	"net/http"

	_ "github.com/lib/pq"
)

func isZap(logger *loggers.Logger) bool {
	if logger == nil {
		return false
	}

	_, ok := (*logger).(loggers.Zap)

	return ok
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

const (
	SERVER_PORT = "8080"
)

func main() {
	var logger loggers.Logger

	zap := loggers.Zap{}
	zap.SetZapLogger()
	logger = zap
	defer syncLogger(&logger)

	var databaseConfig databases.Database
	db, err := sql.Open("postgres", databaseConfig.DBInfo())
	defer db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}

	httpServer := http.Server{
		Addr: fmt.Sprintf(":%s", SERVER_PORT),
	}

	errChannel := make(chan error)
	go func() {
		errChannel <- httpServer.ListenAndServe()
	}()

	serverWelcome := fmt.Sprintf("Server running at port %s", SERVER_PORT)
	logger.Info(serverWelcome)

	for {
		select {
		case err = <-errChannel:
			logger.Fatal(err.Error())
		}
	}
}
