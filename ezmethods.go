package ezsesh

import (
	"github.com/google/uuid"
	"net/http"
)

type EzStoreMethods interface {
	Create(w http.ResponseWriter, assocValue string) error
	GenerateCookieVerifier() (originalVerifier string, hashedVerifier string, err error)
	GenerateCookie(id uuid.UUID, options *EzOptions) (cookieReference *EZCookie, originalVerifier string)
	/*
		Delete()

		OnGenerate()
		OnDelete()
		OnExpire()
	*/
}
