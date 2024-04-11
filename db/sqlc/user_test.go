package db

import (
	"context"
	_ "github.com/lib/pq" // import PostgreSQL driver
	"testing"
)

// TestCreateUser tests the CreateUser function.
func TestCreateUser(t *testing.T) {

	ctx := context.Background()

	// Test data
	testUser := CreateUserParams{
		HashedPassword: "hashed_password",
		Email:          "test1@example.com",
	}

	// Execute the CreateUser function
	user, err := testQueries.CreateUser(ctx, testUser)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// Assert that the user is created with the correct data
	if user.Email != testUser.Email {
		t.Errorf("expected email %s, got %s", testUser.Email, user.Email)
	}
	// Additional assertions can be added for other fields if needed
}

// TestGetUser tests the GetUser function.
func TestGetUser(t *testing.T) {
	ctx := context.Background()

	// Test data
	testEmail := "test1@example.com"

	// Execute the GetUser function
	user, err := testQueries.GetUser(ctx, testEmail)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	// Assert that the user is retrieved with the correct email
	if user.Email != testEmail {
		t.Errorf("expected email %s, got %s", testEmail, user.Email)
	}
	// Additional assertions can be added for other fields if needed
}
