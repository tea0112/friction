package main

import (
	"database/sql"
	"fmt"
	"friction/dbs"
	"friction/handlers"
	"friction/loggers"
	"net/http"

	_ "github.com/lib/pq"
)

func isZap(logger loggers.Logger) bool {
	_, ok := (logger).(*loggers.Zap)

	return ok
}

func syncLogger(logger loggers.Logger) {
	if isZap(logger) == false {
		return
	}
	zap := (logger).(*loggers.Zap)
	zap.ZapLogger.Sync()
}

const (
	SERVER_PORT = "8080"
)

func main() {
	var logger loggers.Logger

	zap := loggers.Zap{}
	zap.SetZapLogger()
	logger = &zap
	defer syncLogger(logger)

	db, err := sql.Open("postgres", dbs.DBInfo())
	defer db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}

	mux := handlers.SetupHandlers(db, logger)

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", SERVER_PORT),
		Handler: mux,
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
