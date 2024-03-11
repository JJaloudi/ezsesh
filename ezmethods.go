package ezsesh

import (
	"net/http"
)

type EzStoreMethods interface {
	Create(w http.ResponseWriter, assocValue string) error

	GetByAssociation(assocValue string) (session interface{}, err error)
	GetSessionByID(sessionId string, destination interface{}) error

	DeleteSession(sessionId string) error
	DeleteSessionByAssoc(assoc string) error
}
