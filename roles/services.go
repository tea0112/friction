package roles

import (
	"context"
	"database/sql"
	"friction/loggers"
)

type Service interface {
	RetrieveAllRoles(context.Context) ([]Role, error)
	SaveRole(context.Context, Role) (Role, error)
	RetrieveRole(context.Context, int64) (Role, error)
}

type ServiceImpl struct {
	DB     *sql.DB
	Logger loggers.Logger
	R      Repository
}

func NewService(db *sql.DB, logger loggers.Logger) Service {
	r := NewRepository(db, logger)
	return ServiceImpl{DB: db, Logger: logger, R: r}
}

func (s ServiceImpl) RetrieveAllRoles(ctx context.Context) ([]Role, error) {
	roles, err := s.R.FindAllRoles(ctx)
	if err != nil {
		s.Logger.Error(err.Error())
		return roles, err
	}

	return roles, nil
}

func (s ServiceImpl) SaveRole(ctx context.Context, role Role) (Role, error) {
	role, err := s.R.PersistRole(ctx, role)
	if err != nil {
		s.Logger.Error(err.Error())
		return role, err
	}

	return role, nil
}

func (s ServiceImpl) RetrieveRole(ctx context.Context, id int64) (Role, error) {
	role, err := s.R.FindById(ctx, id)
	if err != nil {
		return role, err
	}

	return role, nil
}
