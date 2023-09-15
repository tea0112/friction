package handlers

import (
	"friction/users"
	"net/http"
)

func init() {
	http.Handle("/users", new(users.UserHandler))
}
