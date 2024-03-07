package ezsesh

import "net/http"

type EzOptions struct {
	Table       string
	Association string
	CookieName  string
	Lifetime    int64

	HttpOnly bool
	Secure   bool
	SameSite http.SameSite

	SingleToken bool
}
