package stores

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jjaloudi/ezsesh"
	"github.com/jmoiron/sqlx"
	"net/http"
	"reflect"
)

type EzSqlxStore struct {
	Options *ezsesh.EzOptions
	db      *sqlx.DB
}

func CreateEZSqlxStore(options *ezsesh.EzOptions, database *sqlx.DB) *EzSqlxStore {
	return &EzSqlxStore{
		db:      database,
		Options: options,
	}
}

func (store *EzSqlxStore) Create(w http.ResponseWriter, assoc string) error {
	uid := uuid.New()
	cookie, original := ezsesh.GenerateCookie(store.Options, ezsesh.StripUUID(uid.String()))

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

func (store *EzSqlxStore) GetByAssociation(assocValue string) (session interface{}, err error) {
	query := fmt.Sprintf("select * from %s where %s = $1", store.Options.Table, store.Options.Association)
	err = store.db.Select(session, query, assocValue)

	return session, err
}

func (store *EzSqlxStore) GetSessionByID(sessionId string, destination interface{}) error {
	query := fmt.Sprintf("select * from %s where session_id = $1", store.Options.Table)

	destPtr := reflect.ValueOf(destination)
	if err := store.db.QueryRowx(query, sessionId).StructScan(destPtr.Interface()); err != nil {
		return nil
	}

	return nil
}

func (store *EzSqlxStore) DeleteSession(sessionId string) error {
	return nil
}
func (store *EzSqlxStore) DeleteSessionByAssoc(assoc string) error {
	return nil
}
