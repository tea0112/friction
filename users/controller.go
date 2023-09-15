package users

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func (UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
}
