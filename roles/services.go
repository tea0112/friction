package roles

import (
	"database/sql"
	"friction/loggers"
)

type Service interface {
	RetrieveAllRoles() ([]Role, error)
	SaveRole(Role) (Role, error)
	RetrieveRole(int64) (Role, error)
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

func (s ServiceImpl) RetrieveAllRoles() ([]Role, error) {
	roles, err := s.R.FindAllRoles()
	if err != nil {
		s.Logger.Error(err.Error())
		return roles, err
	}

	return roles, nil
}

func (s ServiceImpl) SaveRole(role Role) (Role, error) {
	role, err := s.R.PersistRole(role)
	if err != nil {
		s.Logger.Error(err.Error())
		return role, err
	}

	return role, nil
}

func (s ServiceImpl) RetrieveRole(id int64) (Role, error) {
	role, err := s.R.FindById(id)
	if err != nil {
		return role, err
	}

	return role, nil
}
