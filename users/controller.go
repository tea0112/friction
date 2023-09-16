package users

import (
	"database/sql"
	"fmt"
	"friction/loggers"
	"net/http"
)

type UserHandler struct {
	DB     *sql.DB
	Logger loggers.Logger
}

func (UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
}
