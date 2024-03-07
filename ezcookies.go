package ezsesh

import (
	"github.com/google/uuid"
	"net/http"
)

type EZCookie struct {
	Cookie *http.Cookie
	ID     uuid.UUID
}
