package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host    = "localhost"
	port    = 5432
	user   = "postgres"
	password = "postoGureiy15"
	dbname = "simple_book"
)

var DB *sql.DB

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
