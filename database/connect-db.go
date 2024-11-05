package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=todolist sslmode=disable password=postgres host=localhost"
	db, err := sql.Open("postgres", connStr)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// DB = db
	// defer db.Close()
	return db, err
}
