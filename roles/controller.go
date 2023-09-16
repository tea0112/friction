package roles

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"friction/loggers"
	"friction/utils"
	"io"
	"net/http"
)

type HandlerImpl struct {
	DB     *sql.DB
	Logger loggers.Logger
	s      Service
}

func NewHandler(db *sql.DB, logger loggers.Logger) http.Handler {
	s := NewService(db, logger)
	return HandlerImpl{DB: db, Logger: logger, s: s}
}

func (h HandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.saveRole(w, r)
	}
}

// TODO handle the uri
func (h HandlerImpl) get(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathLen := len(path)

	rolesPath := "/roles"
	rolesPathLen := len(rolesPath)

	if path == rolesPath {
		h.getRoles(w, r)
	} else if pathLen > rolesPathLen {
		roleIdPath := path[rolesPathLen:]
		id, err := utils.GetIdFromPath(roleIdPath)
		if err != nil {
			responseNotFound(w, err, h.Logger)
			return
		}
		h.GetRole(w, r, id)
	}
}

func (h HandlerImpl) getRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	roles, err := h.s.RetrieveAllRoles()
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(roles)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = w.Write(payload.Bytes())
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func (h HandlerImpl) saveRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var role Role
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&role)
	if err != io.EOF && err != nil {
		responseNotFound(w, err, h.Logger)
		return
	}

	role, err = h.s.SaveRole(role)
	if err != nil {
		responseNotFound(w, err, h.Logger)
		return
	}

	var responsePayload bytes.Buffer
	encoder := json.NewEncoder(&responsePayload)

	err = encoder.Encode(role)
	if err != nil {
		responseNotFound(w, err, h.Logger)
		return
	}

	w.Write(responsePayload.Bytes())
}

func (h HandlerImpl) GetRole(w http.ResponseWriter, r *http.Request, id int64) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	role, err := h.s.RetrieveRole(id)
	if err != nil {
		h.Logger.Error(err.Error())
		responseNotFound(w, err, h.Logger)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(role)
	if err != nil {
		h.Logger.Error(err.Error())
		responseNotFound(w, err, h.Logger)
		return
	}

	w.Write(payload.Bytes())
}

func responseNotFound(w http.ResponseWriter, err error, logger loggers.Logger) {
	logger.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
