package permissions

type Role interface {
	String() string
}

type Permission struct {
	id    int64
	name  string
	roles []Role
}
