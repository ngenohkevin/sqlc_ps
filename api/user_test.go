package api

import (
	"fmt"
	"github.com/golang/mock/gomock"
	mockdb "github.com/ngenohkevin/sqlc_ps/db/mock"
	db "github.com/ngenohkevin/sqlc_ps/db/sqlc"
	"github.com/ngenohkevin/sqlc_ps/util"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).
		Return(user, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%d", user.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, req)
}

func randomUser() db.User {
	return db.User{
		ID:        util.RandomInt(1, 1000),
		Firstname: util.RandomString(6),
		Lastname:  util.RandomString(8),
	}
}
