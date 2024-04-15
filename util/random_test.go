package util

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	// Test case: generate a random string of size 10
	size := 10
	randomStr := RandomString(size)

	// Check if the length of the generated string is equal to the specified size
	if len(randomStr) != size {
		t.Errorf("Generated string length is incorrect. Expected: %d, Got: %d", size, len(randomStr))
	}

	// Check if the generated string contains only characters from the "abc" constant
	match, _ := regexp.MatchString("^[a-zA-Z]+$", randomStr)
	if !match {
		t.Errorf("Generated string contains invalid characters: %s", randomStr)
	}

	newRandomStr := RandomString(size)
	if randomStr == newRandomStr {
		t.Errorf("Two generated strings are the same: %s", randomStr)
	}
}

func TestRandomEmail(t *testing.T) {
	// Test case: generate a random email
	randomEmail := RandomEmail()

	// Check if the generated email contains "@" symbol
	if !strings.Contains(randomEmail, "@") {
		t.Errorf("Generated email is invalid: %s", randomEmail)
	}
}
