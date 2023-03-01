package db

import (
	"context"
	"github.com/ngenohkevin/sqlc_ps/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTodo(t *testing.T, user User) Todo {
	arg := CreateTodoParams{
		UsersID: user.ID,
		Task:    util.RandomTasks(),
		Done:    util.RandomBool(),
	}
	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.UsersID, todo.UsersID)
	require.Equal(t, arg.Task, todo.Task)
	require.Equal(t, arg.Done, todo.Done)

	require.NotEmpty(t, todo.ID)

	return todo
}

func TestCreateTodo(t *testing.T) {
	todo := createRandomUser(t)
	createRandomTodo(t, todo)
}

func TestListTodo(t *testing.T) {

}
