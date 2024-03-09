package ezsesh

import (
	"github.com/google/uuid"
	"net/http"
)

type EzStoreMethods interface {
	Create(w http.ResponseWriter, assocValue string) error
	GenerateCookieVerifier() (originalVerifier string, hashedVerifier string, err error)
	GenerateCookie(id uuid.UUID) (cookieReference *EZCookie, originalVerifier string)

	GetUserSession(assocValue string) (session interface{}, err error)
	GetSessionByID(sessionId string) error

	DeleteSession(sessionId string) error
	DeleteSessionByAssoc(assoc string) error
}
