package ezsesh

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type EzStore struct {
	EzStoreMethods
	Options *EzOptions
}

func (store *EzStore) Create(w http.ResponseWriter, assocValue string) error {
	return errors.New("using the default store, this has nothing implemented yet")
}

func (store *EzStore) GenerateCookieVerifier() (string, string, error) {
	cBytes := make([]byte, 16)
	originalBytes := make([]byte, len(cBytes))

	_, err := rand.Read(cBytes)
	if err != nil {
		return "", "", err
	}

	copy(originalBytes, cBytes)

	hash := sha256.New()
	hash.Write(cBytes)

	return hex.EncodeToString(originalBytes), hex.EncodeToString(hash.Sum(nil)), nil
}

func (store *EzStore) GenerateCookie(id uuid.UUID, options *EzOptions) (*EZCookie, string) {
	original, verifier, err := store.GenerateCookieVerifier()
	if err != nil {
		return nil, ""
	}

	cookie := &http.Cookie{
		Name:    options.CookieName,
		Value:   hex.EncodeToString([]byte(id.String())) + "-" + verifier,
		Expires: time.Now().Add(time.Duration(options.Lifetime) * time.Minute),

		HttpOnly: options.HttpOnly,
		Secure:   options.Secure,
		SameSite: options.SameSite,
	}

	return &EZCookie{
		Cookie: cookie,
		ID:     id,
	}, original
}
