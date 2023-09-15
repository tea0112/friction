package main

import (
	"database/sql"
	"fmt"
	"friction/databases"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var databaseConfig databases.Database
	conn, err := sql.Open("postgres", databaseConfig.DBInfo())
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query("SELECT version();")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var ver string
		rows.Scan(&ver)
		fmt.Println(ver)
	}

	rows.Close()

	conn.Close()
}
