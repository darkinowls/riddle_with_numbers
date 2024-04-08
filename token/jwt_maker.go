package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (ITokenMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	tkn, err := jwt.ParseWithClaims(token, payload, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalidToken
		}
		// bytes
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		// check if expired
		tErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(tErr.Inner, errTokenExpired) {
			return nil, errTokenExpired
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, errInvalidToken
	}
	return payload, nil
}
