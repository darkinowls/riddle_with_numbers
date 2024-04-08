package token

import (
	"github.com/stretchr/testify/require"
	"riddle_with_numbers/util"
	"testing"
	"time"
)

func TestPaseto(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)

	username := util.RandomOwner()
	dur := time.Minute

	token, err := maker.CreateToken(username, dur)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.Equal(t, username, payload.Username)
	require.NotZero(t, payload.Id)

	issued := time.Now()
	expired := issued.Add(dur)

	require.WithinDuration(t, payload.IssuedAt, issued, time.Second)
	require.WithinDuration(t, payload.ExpiredAt, expired, time.Second)
}

func TestExpiredTokenPaseto(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errTokenExpired.Error())
	require.Nil(t, payload)
}
