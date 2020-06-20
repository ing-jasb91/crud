package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// getConnection obtiene una conexión a la BD
func getConnection() *sql.DB {
	dsn := "postgres://golang:golang@localhost:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
