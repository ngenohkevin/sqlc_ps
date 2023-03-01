package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	db2 "github.com/ngenohkevin/sqlc_ps/db/sqlc"
	"log"
)

func main() {
	conn, err := sql.Open("postgres", "postgresql://postgres:dylstar20@localhost:5432/sqlc_db?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}
	db := db2.New(conn)

	user, err := db.CreateUser(context.Background(), db2.CreateUserParams{
		Firstname: "Jay",
		Lastname:  "Maina",
	})

	todo, err := db.CreateTodo(context.Background(), db2.CreateTodoParams{
		UsersID: user.ID,
		Task:    "Bathing",
		Done:    true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(user, todo)

}
