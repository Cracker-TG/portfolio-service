package helpers

import (
	"testing"
)

func TestHashAndDecryText(t *testing.T) {
	text := "ONEMORE!"
	var h HelpersInteface = new(Helpers)

	_, err := h.Encryption(text)

	if err != nil {
		t.Error("encryption error", err)
		return
	}

	check := h.Decryption("ENCRYPTION", text)

	if !check {
		t.Error("Text Not Match!", "digest ")
		return
	}
}
