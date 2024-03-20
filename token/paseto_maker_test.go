package token

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetPayloadFromToken(t *testing.T) {
	issued_at := time.Now()
	expired_at := issued_at.Add(365 * 24 * time.Hour)

	tokenID, err := uuid.NewRandom()
	require.NoError(t, err)

	// Create a Paseto token with test data
	token := paseto.NewToken()
	token.Set("id", tokenID.String())
	token.Set("username", "testuser")
	token.SetIssuedAt(issued_at) // August 13, 2021 00:00:00 UTC
	token.SetExpiration(expired_at) // August 13, 2022 00:00:00 UTC

	// Call the function under test
	payload, err := getPayloadFromToken(&token)
	require.NoError(t, err)

	// Assert the expected values
	require.Equal(t, tokenID, payload.ID) // Remove dereference operator
	require.Equal(t, "testuser", payload.Username)
	require.Equal(t, issued_at.Truncate(time.Second), payload.IssuedAt.Truncate(time.Second))
	require.Equal(t, expired_at.Truncate(time.Second), payload.ExpiredAt.Truncate(time.Second))
}