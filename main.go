package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ngenohkevin/sqlc_ps/api"
	db "github.com/ngenohkevin/sqlc_ps/db/sqlc"
	"github.com/ngenohkevin/sqlc_ps/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
