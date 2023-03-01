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
	user := createRandomUser(t)

	for i := 0; i < 10; i++ {
		createRandomTodo(t, user)
	}

	arg := ListTodoParams{
		UsersID: user.ID,
		Limit:   5,
		Offset:  5,
	}
	todos, err := testQueries.ListTodo(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, todos, 5)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
		require.Equal(t, arg.UsersID, todo.UsersID)
	}

}
