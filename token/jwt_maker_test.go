package token

import (
	"riddle_with_numbers/util"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJwt(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)

	email := util.RandomEmail()
	dur := time.Minute

	token, err := maker.CreateToken(email, dur)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, email, payload.Email)
	require.NotZero(t, payload.Id)

	issued := time.Now()
	expired := issued.Add(dur)

	require.WithinDuration(t, payload.IssuedAt, issued, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, expired, time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)

	token, err := maker.CreateToken(util.RandomEmail(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errTokenExpired.Error())
	require.Nil(t, payload)
}

// REMOVE THIS. USELESS
func TestTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomEmail(), time.Minute)
	require.NoError(t, err)

	tkn := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	// alg=None
	tokenStr, err := tkn.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
	maker, _ := NewJWTMaker(util.RandomString(32))

	payload, err = maker.VerifyToken(tokenStr)
	require.Error(t, err)
	require.EqualError(t, err, errInvalidToken.Error())
	require.Nil(t, payload)
}
