package roles

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"friction/ctxkeys"
	"friction/loggers"
	"io"
	"net/http"
	"strconv"
)

type Controller interface {
	GetRoles(http.ResponseWriter, *http.Request)
	GetRole(http.ResponseWriter, *http.Request)
	SaveRole(http.ResponseWriter, *http.Request)
}

type ControllerImpl struct {
	db          *sql.DB
	logger      loggers.Logger
	roleService Service
}

func NewController(db *sql.DB, logger loggers.Logger) Controller {
	return ControllerImpl{
		db:          db,
		logger:      logger,
		roleService: NewService(db, logger),
	}
}

func (c ControllerImpl) GetRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	roles, err := c.roleService.RetrieveAllRoles(ctx)
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(roles)
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = w.Write(payload.Bytes())
	if err != nil {
		c.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func (c ControllerImpl) SaveRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var role Role
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&role)
	if err != io.EOF && err != nil {
		responseNotFound(w, err, c.logger)
		return
	}

	role, err = c.roleService.SaveRole(ctx, role)
	if err != nil {
		responseNotFound(w, err, c.logger)
		return
	}

	var responsePayload bytes.Buffer
	encoder := json.NewEncoder(&responsePayload)

	err = encoder.Encode(role)
	if err != nil {
		responseNotFound(w, err, c.logger)
		return
	}

	w.Write(responsePayload.Bytes())
}

func (c ControllerImpl) GetRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := ctx.Value(ctxkeys.ParamsCtxKey{}).([]string)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		c.logger.Error(err.Error())
		responseNotFound(w, err, c.logger)
		return
	}

	role, err := c.roleService.RetrieveRole(ctx, id)
	if err != nil {
		c.logger.Error(err.Error())
		responseNotFound(w, err, c.logger)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(role)
	if err != nil {
		c.logger.Error(err.Error())
		responseNotFound(w, err, c.logger)
		return
	}

	w.Write(payload.Bytes())
}

func responseNotFound(w http.ResponseWriter, err error, logger loggers.Logger) {
	logger.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
