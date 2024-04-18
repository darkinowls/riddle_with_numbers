package db

import (
	"context"
	_ "github.com/lib/pq" // import PostgreSQL driver
	"riddle_with_numbers/util"
	"testing"
)

func TestUserQueries(t *testing.T) {

	email := util.RandomEmail()

	testCases := []struct {
		name          string
		checkResponse func(t *testing.T)
	}{
		{
			name: "TestCreateUser",
			checkResponse: func(t *testing.T) {
				testUser := CreateUserParams{
					HashedPassword: "hashed_password",
					Email:          email,
				}

				// Execute the CreateUser function
				user, err := testQueries.CreateUser(context.Background(), testUser)
				if err != nil {
					t.Fatalf("CreateUser failed: %v", err)
				}

				// Assert that the user is created with the correct data
				if user.Email != testUser.Email {
					t.Errorf("expected email %s, got %s", testUser.Email, user.Email)
				}
			},
		},
		{
			name: "TestGetUser",
			checkResponse: func(t *testing.T) {
				testEmail := email

				// Execute the GetUser function
				user, err := testQueries.GetUser(context.Background(), testEmail)
				if err != nil {
					t.Fatalf("GetUser failed: %v", err)
				}

				// Assert that the user is retrieved with the correct email
				if user.Email != testEmail {
					t.Errorf("expected email %s, got %s", testEmail, user.Email)
				}
				if user.HashedPassword != "hashed_password" {
					t.Errorf("expected hashed password %s, got %s", "hashed_password", user.HashedPassword)

				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkResponse(t)
		})
	}

}
