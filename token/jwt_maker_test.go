package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zzooman/zapp-server/utils"
)

func TestJWTMaker_VerifyToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomString(6)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.Username)
	require.NotZero(t, payload.ID)	
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)	
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)
	
	duration := -time.Minute

	token, err := maker.CreateToken(utils.RandomString(6), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}

func TestInvalideJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err := maker.VerifyToken("invalid-token")
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}