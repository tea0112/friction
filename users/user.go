package users

import "friction/roles"

type User struct {
	id       int64
	email    string
	username string
	password string
	roles    []roles.Role
}

func (User) String() string {
	return "users"
}
