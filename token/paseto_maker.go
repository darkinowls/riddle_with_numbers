package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto    *paseto.V2
	secretKey []byte
}

func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.secretKey, payload, nil)
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.secretKey, payload, nil)
	if err != nil {
		return nil, errInvalidToken

	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func NewPasetoMaker(secretKey string) (ITokenMaker, error) {
	if len(secretKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", minSecretKeySize)
	}

	maker := &PasetoMaker{
		paseto:    paseto.NewV2(),
		secretKey: []byte(secretKey),
	}
	return maker, nil
}
