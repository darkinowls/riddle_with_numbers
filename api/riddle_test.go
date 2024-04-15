package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"riddle_with_numbers/riddle"
	"riddle_with_numbers/token"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func addAuthorizationHeader(
	t *testing.T,
	request *http.Request,
	tokenMaker token.ITokenMaker,
	email string,
	duration time.Duration,
) {
	tkn, err := tokenMaker.CreateToken(email, duration)
	require.NoError(t, err)
	authHeader := "Bearer " + tkn
	request.Header.Set(authorizationHeaderKey, authHeader)
}

/////////////////////

func TestCheckIfAlive(t *testing.T) {
	server, err := NewTestServer()
	assert.NoError(t, err)
	router := server.router

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestSolveRiddle(t *testing.T) {
	server, err := NewTestServer()
	assert.NoError(t, err)
	router := server.router

	matrix := [][]int{
		{4, 2, 4, 8}, {8, 6, 6, 8}, {4, 2, 6, 6}, {2, 2, 6, 6},
	}
	jsonMatrix, _ := json.Marshal(matrix)

	testCases := []struct {
		name          string
		checkResponse func(t *testing.T)
	}{
		{
			name: "401 check auth",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusUnauthorized, w.Code)
			},
		},

		{
			name: "401 invalid token",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
				authHeader := "Bearer " + "invalid_token_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
				req.Header.Set(authorizationHeaderKey, authHeader)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusUnauthorized, w.Code)
			},
		},

		{
			name: "200 create",
			checkResponse: func(t *testing.T) {
				req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
				addAuthorizationHeader(t, req, server.tokenMaker, "user1@example.com", time.Minute)
				assert.NoError(t, err)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)
				count, err := strconv.Atoi(w.Body.String())
				if err != nil {
					return
				}
				assert.True(t, count == 1)
			},
		},
		{
			name: "200 create and get solution",
			checkResponse: func(t *testing.T) {
				req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
				addAuthorizationHeader(t, req, server.tokenMaker, "user1@example.com", time.Minute)
				assert.NoError(t, err)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)
				count, err := strconv.Atoi(w.Body.String())
				if err != nil {
					return
				}
				assert.True(t, count == 1)

				// Get solution

				w = httptest.NewRecorder()
				req, err = http.NewRequest("GET", "/solution", bytes.NewReader(jsonMatrix))
				addAuthorizationHeader(t, req, server.tokenMaker, "user1@example.com", time.Minute)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)
				var result [][]riddle.Cell

				println(w.Body.String())

				err = json.Unmarshal(w.Body.Bytes(), &result)
				assert.NoError(t, err)
				expected := riddle.GetExampleResult()
				assert.True(t, riddle.CompareMatrices(expected, result))
			},
		},
		{
			name: "Error solve riddle 400 and 404",
			checkResponse: func(t *testing.T) {
				matrix = [][]int{
					{4, 2, 4, 8},
				}
				jsonMatrix, _ = json.Marshal(matrix)

				w := httptest.NewRecorder()
				req, err := http.NewRequest("POST", "/solve", bytes.NewReader(jsonMatrix))
				addAuthorizationHeader(t, req, server.tokenMaker, "user1@example.com", time.Minute)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)

				assert.Equal(t, http.StatusBadRequest, w.Code)

				w = httptest.NewRecorder()
				req, err = http.NewRequest("GET", "/solution", bytes.NewReader(jsonMatrix))
				addAuthorizationHeader(t, req, server.tokenMaker, "user1@example.com", time.Minute)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusNotFound, w.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkResponse(t)
		})
	}

}

func TestCondition(t *testing.T) {
	server, err := NewTestServer()
	assert.NoError(t, err)
	router := server.router

	testCases := []struct {
		name          string
		checkResponse func(t *testing.T)
	}{
		{
			name: "Error num input 400",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("POST", "/generate/1000", nil)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
		{
			name: "Generate conditions",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("POST", "/generate/2", nil)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)
			},
		},
		{
			name: "Check condition",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/condition/1", nil)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code)

				// Read the response body as bytes
				bodyBytes := w.Body.Bytes()
				var matrix [][]int
				err = json.Unmarshal(bodyBytes, &matrix)
				assert.NoError(t, err)
				assert.Equal(t, 2, len(matrix)) // Assuming the matrix is of size 3x3

			},
		},

		{
			name: "400 bad id",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/condition/aaa", nil)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},

		{
			name: "404",
			checkResponse: func(t *testing.T) {
				w := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/condition/100", nil)
				assert.NoError(t, err)
				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusNotFound, w.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkResponse(t)
		})
	}
}
