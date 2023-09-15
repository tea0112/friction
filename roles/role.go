package roles

import "friction/permissions"

type User interface {
	String() string
}

type Role struct {
	id          int64
	name        string
	users       []User
	permissions []permissions.Permission
}

func (Role) String() string {
	return "roles"
}
