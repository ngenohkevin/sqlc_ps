// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import ()

type Todo struct {
	ID      int64  `json:"id"`
	UsersID int64  `json:"users_id"`
	Task    string `json:"task"`
	Done    bool   `json:"done"`
}

type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
