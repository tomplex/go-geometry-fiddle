package main

import (
	"fmt"
	"database/sql"
)

var connection *sql.DB

func init() {
	connectionString := "host=172.17.0.2 dbname=aqueousband_com user=postgres password=postgres port=5432 sslmode=disable"

	var err error
	connection, err = sql.Open(
		"postgres",
		connectionString,
	)

	err = connection.Ping()
	check(err)
}

func TestConnect() {
	var pid int

	err := connection.QueryRow("SELECT pg_backend_pid()").Scan(&pid)
	check(err)
	fmt.Printf("pid: %v\n\n\n", pid)
}
