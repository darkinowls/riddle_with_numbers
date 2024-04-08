package token

import (
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errTokenExpired
	}
	return nil
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	payload := &Payload{
		Id:        token,
		Email:     email,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}
	return payload, nil

}
