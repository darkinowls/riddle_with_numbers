package token

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPayload_Valid(t *testing.T) {
	// Test case: token is valid
	now := time.Now()
	expiredAt := now.Add(time.Hour)
	payload := &Payload{
		ExpiredAt: expiredAt,
	}
	if err := payload.Valid(); err != nil {
		t.Errorf("Expected token to be valid, got error: %v", err)
	}

	// Test case: token is expired
	expiredAt = now.Add(-time.Hour) // Set expiration time in the past
	payload = &Payload{
		ExpiredAt: expiredAt,
	}
	expectedError := errTokenExpired
	if err := payload.Valid(); err != expectedError {
		t.Errorf("Expected token to be expired with error: %v, got error: %v", expectedError, err)
	}
}

func TestNewPayload(t *testing.T) {
	// Test case: create payload with valid parameters
	email := "test@example.com"
	duration := time.Hour
	payload, err := NewPayload(email, duration)
	if err != nil {
		t.Errorf("Failed to create new payload: %v", err)
	}

	// Check if payload fields are correctly set
	if payload.Email != email {
		t.Errorf("Expected email %s, got %s", email, payload.Email)
	}
	now := time.Now()

	require.WithinDurationf(t, payload.IssuedAt, now, time.Second, "Invalid creation payload timestamps, issued at: %v, now: %v", payload.IssuedAt, now)

	expectedExpiredAt := now.Add(duration)

	require.WithinDurationf(t, payload.ExpiredAt, expectedExpiredAt, time.Second, "Invalid expire payload timestamps, expired at: %v, expected: %v", payload.ExpiredAt, expectedExpiredAt)

}
