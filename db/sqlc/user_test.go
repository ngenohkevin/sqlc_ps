package db

import (
	"context"
	"database/sql"
	"github.com/ngenohkevin/sqlc_ps/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Firstname: util.RandomString(5),
		Lastname:  util.RandomString(6),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Firstname, user.Firstname)
	require.Equal(t, arg.Lastname, user.Lastname)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Firstname, user2.Firstname)
	require.Equal(t, user1.Lastname, user2.Lastname)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)

}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	user, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, user, 5)

	for _, users := range user {
		require.NotEmpty(t, users)
	}
}
