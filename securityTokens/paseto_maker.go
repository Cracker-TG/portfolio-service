package securityTokens

import (
	"os"
	"time"

	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasetoInteface interface {
	CreateToken(user_id *primitive.ObjectID, username *string, duration *time.Duration) (string, error)
	VerfifyToken(token *string) (bool, *Payload)
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Payload struct {
	ID        primitive.ObjectID
	Username  string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewPayload(user_id *primitive.ObjectID, username *string, duration *time.Duration) (*Payload, error) {

	payload := &Payload{
		ID:        *user_id,
		Username:  *username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(*duration),
	}

	return payload, nil
}

func (maker PasetoMaker) CreateToken(user_id *primitive.ObjectID, username *string, duration *time.Duration) (string, error) {
	maker.paseto = paseto.NewV2()
	maker.symmetricKey = []byte(os.Getenv("SymmetricKey"))

	payload, err := NewPayload(user_id, username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker PasetoMaker) VerfifyToken(token *string) (bool, *Payload) {
	maker.paseto = paseto.NewV2()
	maker.symmetricKey = []byte(os.Getenv("SymmetricKey"))
	payload := Payload{}
	err := maker.paseto.Decrypt(*token, maker.symmetricKey, &payload, nil)

	if err != nil {
		return false, nil
	}
	/** expri **/
	if ValidExprired(&payload) {
		return false, nil
	}

	return true, &payload
}

func ValidExprired(payload *Payload) bool {
	if time.Now().After(payload.ExpiredAt) {
		return true
	}
	return false
}
