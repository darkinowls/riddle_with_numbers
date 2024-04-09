package token

import (
	"fmt"
	"time"
)

var errInvalidToken = fmt.Errorf("invalid token")
var errTokenExpired = fmt.Errorf("token is expired")

// factory pattern
type ITokenMaker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
