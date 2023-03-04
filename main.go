package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ngenohkevin/sqlc_ps/api"
	db "github.com/ngenohkevin/sqlc_ps/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
