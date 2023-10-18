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

const (
	SERVER_PORT = "8080"
)

func main() {
	logger := loggers.NewLogger()
	defer loggers.SyncLogger(logger)

	db, err := sql.Open("postgres", dbs.DBInfo())
	defer db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		logger.Fatal(err.Error())
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", SERVER_PORT),
		Handler: handlers.InitMux(db, logger),
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
