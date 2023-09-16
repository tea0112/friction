package roles

import "friction/permissions"

type User interface {
	String() string
}

type Role struct {
	Id          int64                    `json:"id"`
	Name        string                   `json:"name"`
	Users       []User                   `json:"users,omitempty"`
	Permissions []permissions.Permission `json:"permissions,omitempty"`
}

func (Role) String() string {
	return "roles"
}
