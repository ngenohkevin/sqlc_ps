package main

import (
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable"
)

func main() {

}
