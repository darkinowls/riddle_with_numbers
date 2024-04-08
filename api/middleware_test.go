package api

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/require"
//	"net/http"
//	"net/http/httptest"
//	"riddle_with_numbers/token"
//	"testing"
//	"time"
//)
//
//func addAuthorizationHeader(
//	t *testing.T,
//	request *http.Request,
//	tokenMaker token.ITokenMaker,
//	username string,
//	duration time.Duration,
//) {
//	tkn, err := tokenMaker.CreateToken(username, duration)
//	require.NoError(t, err)
//	authHeader := "Bearer " + tkn
//	request.Header.Set(authorizationHeaderKey, authHeader)
//}
//
//func TestAuthMiddleware(t *testing.T) {
//	testCases := []struct {
//		name          string
//		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.ITokenMaker)
//		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
//	}{
//		{
//			name: "OK",
//			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.ITokenMaker) {
//				addAuthorizationHeader(t, request, tokenMaker, "user1", time.Minute)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusOK, recorder.Code)
//			},
//		},
//		{
//			name: "401 Not auth",
//			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.ITokenMaker) {
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusUnauthorized, recorder.Code)
//			},
//		},
//		{
//			name: "401 Expired",
//			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.ITokenMaker) {
//				addAuthorizationHeader(t, request, tokenMaker, "user1", -time.Minute)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusUnauthorized, recorder.Code)
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			server := newTestServer(t, nil)
//
//			// create
//			server.router.GET("/auth", authMiddleware(server.tokenMaker), func(c *gin.Context) {
//				c.JSON(http.StatusOK, gin.H{})
//			})
//
//			recorder := httptest.NewRecorder()
//			req, err := http.NewRequest(http.MethodGet, "/auth", nil)
//			require.NoError(t, err)
//
//			tc.setupAuth(t, req, server.tokenMaker)
//			server.router.ServeHTTP(recorder, req)
//			tc.checkResponse(t, recorder)
//
//		})
//	}
//}
