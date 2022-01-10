package helpers

import (
	"crypto/sha512"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type HelpersInteface interface {
	Encryption(text string) (string, error)
	Decryption(digest string, text string) bool
}

type Helpers struct{}

func (h Helpers) Encryption(text string) (string, error) {
	bs := hashSalt(text)
	cost := 10
	bytes, err := bcrypt.GenerateFromPassword(bs, cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (h Helpers) Decryption(digest string, text string) bool {
	bs := hashSalt(text)
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(bs))

	if err != nil {
		return false
	}

	return true
}

func (h Helpers) RandomDigitAndStringEmailConfirm() string {
	return String(6)
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return StringWithCharset(length, charset)
}

func hashSalt(text string) []byte {
	hasher := sha512.New()
	hasher.Write([]byte(text))
	bs := hasher.Sum(nil)
	return bs
}
