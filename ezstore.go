package ezsesh

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type EzStore struct {
	EzStoreMethods
	options *EzOptions
	db      *sqlx.DB
}

func CreateEZStore(options *EzOptions, database *sqlx.DB) *EzStore {
	return &EzStore{
		options: options,
		db:      database,
	}
}

func (store *EzStore) Generate(w http.ResponseWriter, assoc string) {
	uid := uuid.New()
	cookie, original := GenerateCookie(uid, store.options)

	if store.options.SingleToken {
		query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", store.options.Table, store.options.Association)
		_, err := store.db.Exec(query, assoc)

		if err != nil {
			fmt.Errorf("Failed to delete old session row.", err)
			return
		}
	}

	query := fmt.Sprintf(
		"insert into %s (session_id, %s, verifier,  expires_at) values ($1, $2, $3, $4) returning session_id",
		store.options.Table,
		store.options.Association,
	)

	fmt.Println(query)

	_, err := store.db.Exec(query, uid.String(), assoc, original, cookie.Cookie.Expires)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.SetCookie(w, cookie.Cookie)
}
