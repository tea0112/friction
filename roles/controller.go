package roles

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"friction/loggers"
	"friction/utils"
	"io"
	"net/http"
	"strconv"
)

func getRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value(utils.DBCtxKey{}).(*sql.DB)
	logger := ctx.Value(utils.LoggerCtxKey{}).(loggers.Logger)

	roleService := NewService(db, logger, ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	roles, err := roleService.RetrieveAllRoles()
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(roles)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = w.Write(payload.Bytes())
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}

func saveRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db := ctx.Value(utils.DBCtxKey{}).(*sql.DB)
	logger := ctx.Value(utils.LoggerCtxKey{}).(loggers.Logger)

	roleService := NewService(db, logger, ctx)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var role Role
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&role)
	if err != io.EOF && err != nil {
		responseNotFound(w, err, logger)
		return
	}

	role, err = roleService.SaveRole(role)
	if err != nil {
		responseNotFound(w, err, logger)
		return
	}

	var responsePayload bytes.Buffer
	encoder := json.NewEncoder(&responsePayload)

	err = encoder.Encode(role)
	if err != nil {
		responseNotFound(w, err, logger)
		return
	}

	w.Write(responsePayload.Bytes())
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := ctx.Value(utils.DBCtxKey{}).([]string)
	db := ctx.Value(utils.DBCtxKey{}).(*sql.DB)
	logger := ctx.Value(utils.LoggerCtxKey{}).(loggers.Logger)

	roleService := NewService(db, logger, ctx)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		logger.Error(err.Error())
		responseNotFound(w, err, logger)
		return
	}

	role, err := roleService.RetrieveRole(id)
	if err != nil {
		logger.Error(err.Error())
		responseNotFound(w, err, logger)
		return
	}

	var payload bytes.Buffer
	encoder := json.NewEncoder(&payload)

	err = encoder.Encode(role)
	if err != nil {
		logger.Error(err.Error())
		responseNotFound(w, err, logger)
		return
	}

	w.Write(payload.Bytes())
}

func responseNotFound(w http.ResponseWriter, err error, logger loggers.Logger) {
	logger.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
