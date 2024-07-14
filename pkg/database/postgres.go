package database

import (
	"database/sql"
	"fmt"
	"strconv"
)

func OpenDB() (*sql.DB, error) {
	var db *sql.DB
	port, err := strconv.Atoi("5432")
	if err != nil {
		return db, err
	}

	return sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			"localhost", port, "postgres",
			"logoutis", "unsia_go",
		),
	)
}
