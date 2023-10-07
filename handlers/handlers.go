package handlers

import (
	"context"
	"database/sql"
	"friction/loggers"
	"friction/roles"
	"friction/utils"
	"net/http"
	"regexp"
)

type Route struct {
	Method    string
	PathRegex string
	Handler   http.HandlerFunc
}

var routes = []Route{
	Route{http.MethodGet, "/api/roles", roles.GetRole},
}

func SetupHandlers(db *sql.DB, logger loggers.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		notFound := true

		for _, route := range routes {
			if params := matchPath(route.PathRegex, path); params != nil {
				notFound = false

				ctx := r.Context()
				ctx = context.WithValue(ctx, utils.ParamsCtxKey{}, params)
				ctx = context.WithValue(ctx, utils.DBCtxKey{}, db)
				ctx = context.WithValue(ctx, utils.LoggerCtxKey{}, logger)

				r = r.WithContext(ctx)

				route.Handler(w, r)
			}
		}

		if notFound {
			http.NotFound(w, r)
		}
	})

	return mux
}

func matchPath(rg string, path string) []string {
	return regexp.MustCompile(rg).FindStringSubmatch(path)
}
