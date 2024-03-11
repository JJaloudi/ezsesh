package ezsesh

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"
)

func GenerateCookieVerifier() (string, string, error) {
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

/*
If you plan to store a UUID as the assoc value, make sure to run
StripUUID for assoc string

On recovery, you can run RebuildUUID on the identifier in UnwrapCookie.
*/
func GenerateCookie(options *EzOptions, assoc string) (*EZCookie, string) {
	original, verifier, err := GenerateCookieVerifier()
	if err != nil {
		return nil, ""
	}

	cookie := &http.Cookie{
		Name:    options.CookieName,
		Value:   hex.EncodeToString([]byte(assoc)) + "-" + verifier,
		Expires: time.Now().Add(time.Duration(options.Lifetime) * time.Minute),

		HttpOnly: options.HttpOnly,
		Secure:   options.Secure,
		SameSite: options.SameSite,
	}

	return &EZCookie{
		Cookie: cookie,
		ID:     assoc,
	}, original
}

func UnwrapCookie(cookie string) (identifier string, verifier string, err error) {
	if cookie == "" {
		return "", "", errors.New("Cookie string is empty")
	}

	cookieValue := cookie
	split := strings.Split(cookieValue, "-")

	return split[0], split[1], err
}

// Uses crypto/subtle
func CompareVerifier(cookie string, stored string) (compare bool, err error) {
	hashedOriginal := sha256.Sum256([]byte(cookie))

	storedBytes, err := hex.DecodeString(stored)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(hashedOriginal[:], storedBytes) == 1, err
}

/*
	Helper methods to strip/recover a UUID. This is for those of you that want to use a UUID.String(), since this library is meant to have no dependencies
	outside the native Go libraries.

	For instance, you pass GenerateCookies(options, StripUUID(id)) instead of just the plain ID.
*/

func StripUUID(uuidStr string) string {
	return uuidStr[:8] + uuidStr[9:13] + uuidStr[14:18] + uuidStr[19:23] + uuidStr[24:]
}

func RebuildUUID(uuidStr string) (string, error) {
	if len(uuidStr) != 32 {
		return "", errors.New("invalid UUID string length")
	}

	return uuidStr[:8] + "-" + uuidStr[8:12] + "-" + uuidStr[12:16] + "-" + uuidStr[16:20] + "-" + uuidStr[20:], nil
}
