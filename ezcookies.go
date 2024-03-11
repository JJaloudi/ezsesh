package ezsesh

import (
	"net/http"
)

type EZCookie struct {
	Cookie *http.Cookie
	ID     string
}
