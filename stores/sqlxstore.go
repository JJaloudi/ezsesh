package stores

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jjaloudi/ezsesh"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type EzSqlxStore struct {
	ezsesh.EzStore
	db *sqlx.DB
}

func CreateEZSqlxStore(options *ezsesh.EzOptions, database *sqlx.DB) *EzSqlxStore {
	return &EzSqlxStore{
		EzStore: ezsesh.EzStore{
			Options: options,
		},
	}
}

func (store *EzSqlxStore) Create(w http.ResponseWriter, assoc string) error {
	uid := uuid.New()
	cookie, original := store.EzStore.GenerateCookie(uid, store.Options)

	if store.Options.SingleToken {
		query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", store.Options.Table, store.Options.Association)
		_, err := store.db.Exec(query, assoc)

		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf(
		"insert into %s (session_id, %s, verifier,  expires_at) values ($1, $2, $3, $4) returning session_id",
		store.Options.Table,
		store.Options.Association,
	)

	_, err := store.db.Exec(query, uid.String(), assoc, original, cookie.Cookie.Expires)
	if err != nil {
		return err
	}

	http.SetCookie(w, cookie.Cookie)

	return err
}
