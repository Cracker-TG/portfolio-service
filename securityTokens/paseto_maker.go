package securityTokens

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoInteface interface {
	CreateToken(username *string, duration *time.Duration) (string, error)
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Payload struct {
	ID        uuid.UUID
	Username  string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewPayload(username *string, duration *time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  *username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(*duration),
	}

	return payload, nil
}

func (maker PasetoMaker) CreateToken(username *string, duration *time.Duration) (string, error) {
	maker.paseto = paseto.NewV2()
	maker.symmetricKey = []byte(os.Getenv("SymmetricKey"))

	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}
