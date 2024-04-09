package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	db "riddle_with_numbers/db/sqlc"
	"riddle_with_numbers/util"
	"testing"
)

func randomUser() db.User {
	return db.User{
		HashedPassword: util.RandomString(32),
		Email:          util.RandomEmail(),
	}
}

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x any) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", e.arg, e.arg)
}

func TestCreateUserAPI(t *testing.T) {
	server, err := NewTestServer()
	assert.NoError(t, err)
	user := randomUser()

	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "200 - OK",
			body: gin.H{
				"password": user.HashedPassword,
				"email":    user.Email,
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "wrong email",
			body: gin.H{
				"password": user.HashedPassword,
				"email":    "wrong_email",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		//{
		//	name: "500 - Internal Server Error",
		//	body: gin.H{
		//		"password": user.HashedPassword,
		//		"email":    user.Email,
		//	},
		//
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		//	},
		//},
		{
			name: "400 - Duplicate Email",
			body: gin.H{
				"password": user.HashedPassword,
				"email":    user.Email,
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			url := "/auth/create"

			jsonBody, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBody))
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	var gotUser db.User
	err := json.NewDecoder(body).Decode(&gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
