package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "joewalker"
	dbname = "library"
)

func Init() *sql.DB {
	conStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		panic(err)
	}

	if err := con.Ping(); err != nil {
		panic(err)
	}

	return con
}
