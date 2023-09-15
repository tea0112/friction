package databases

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "friction"
	password = "Postgres!23"
	dbname   = "friction"
)

type Database struct{}

func (Database) DBInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
