// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteTodo(ctx context.Context, usersID int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetTodo(ctx context.Context, id int64) (Todo, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListTodo(ctx context.Context, arg ListTodoParams) ([]Todo, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
}

var _ Querier = (*Queries)(nil)
