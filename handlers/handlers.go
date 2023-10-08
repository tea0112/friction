package handlers

import (
	"context"
	"database/sql"
	"friction/ctxkeys"
	"friction/loggers"
	"net/http"
	"regexp"
)

type Route struct {
	Method    string
	PathRegex string
	Handler   http.HandlerFunc
}

func SetupHandlers(db *sql.DB, logger loggers.Logger, routes []Route) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		notFound := true

		for _, route := range routes {
			if params := matchPath(route.PathRegex, path); params != nil {
				notFound = false

				ctx := r.Context()
				ctx = context.WithValue(ctx, ctxkeys.ParamsCtxKey{}, params)

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
