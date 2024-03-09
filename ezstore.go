package ezsesh

import (
	"errors"
	"net/http"
)

type EzStore struct {
	EzStoreMethods
	Options *EzOptions
}

func (store *EzStore) Create(w http.ResponseWriter, assocValue string) error {
	return errors.New("using the default store, this has nothing implemented yet")
}
