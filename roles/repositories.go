package roles

import (
	"database/sql"
	"fmt"
	"friction/loggers"
)

type Repository interface {
	FindAllRoles() ([]Role, error)
	PersistRole(Role) (Role, error)
	FindById(int64) (Role, error)
}

type RepositoryImpl struct {
	DB     *sql.DB
	Logger loggers.Logger
}

func NewRepository(db *sql.DB, logger loggers.Logger) Repository {
	return RepositoryImpl{DB: db, Logger: logger}
}

func (r RepositoryImpl) FindAllRoles() ([]Role, error) {
	roles := make([]Role, 0)

	q := `
select id, name from friction.roles
	`

	rows, err := r.DB.Query(q)
	if err != nil {
		return roles, err
	}

	for rows.Next() {
		var role Role
		rows.Scan(&role.Id, &role.Name)
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return roles, err
	}

	return roles, nil
}

func (r RepositoryImpl) PersistRole(role Role) (Role, error) {
	result, err := r.DB.Exec("insert into friction.roles(name) values($1)", role.Name)
	if err != nil {
		return role, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return role, err
	}
	if rows != 1 {
		return role, fmt.Errorf("no row affected")
	}

	return role, nil
}


func (r RepositoryImpl) FindById(id int64) (Role, error) {
	q := `
select id, name from roles where id = $1
`
	row := r.DB.QueryRow(q, id)	
	
	var role Role
	if err := row.Scan(&role); err != nil {
		return role, err
	}

	return role, nil
}