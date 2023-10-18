package handlers

import (
	"database/sql"
	"friction/loggers"
	"friction/roles"
	"net/http"

	"github.com/go-chi/chi"
)

func InitMux(db *sql.DB, logger loggers.Logger) http.Handler {
	rolesController := roles.NewController(db, logger)

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/roles", func(r chi.Router) {
			r.Get("/", rolesController.GetRoles)
		})
	})

	return r
}
