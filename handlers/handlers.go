package handlers

import (
	"database/sql"
	"friction/loggers"
	"friction/roles"
	"friction/users"
	"net/http"
)

// TODO handle the uri
func SetupHandlers(db *sql.DB, logger loggers.Logger) {
	userHandler := users.UserHandler{DB: db, Logger: logger}
	roleHandler := roles.NewHandler(db, logger)

	http.Handle("/users", userHandler)
	http.Handle("/roles", roleHandler)
}