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

func TestGetTodo(t *testing.T) {
	user := createRandomUser(t)
	todo1 := createRandomTodo(t, user)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.UsersID, todo2.UsersID)
	require.Equal(t, todo1.Task, todo2.Task)
	require.Equal(t, todo1.Done, todo2.Done)
}

func TestUpdateTodo(t *testing.T) {

	user := createRandomUser(t)
	todo1 := createRandomTodo(t, user)

	arg := UpdateTodoParams{
		ID:   todo1.ID,
		Task: util.RandomTasks(),
		Done: util.RandomBool(),
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.UsersID, todo2.UsersID)
	require.Equal(t, arg.Task, todo2.Task)
	require.Equal(t, arg.Done, todo2.Done)

}

//func TestDeleteTodo(t *testing.T) {
//	user := createRandomUser(t)
//	todo1 := createRandomTodo(t, user)
//
//	err := testQueries.DeleteTodo(context.Background(), todo1.UsersID)
//	require.NoError(t, err)
//	//fmt.Println("Delete todo Error: ", err)
//
//	todo2, err := testQueries.GetTodo(context.Background(), todo1.UsersID)
//	require.Error(t, err)
//	//fmt.Println("Get Todo Error: ", err)
//
//	require.EqualError(t, err, sql.ErrNoRows.Error())
//	require.Empty(t, todo2)
//
//}
