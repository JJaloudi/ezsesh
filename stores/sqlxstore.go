package stores

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/jjaloudi/ezsesh"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
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
	cookie, original := store.GenerateCookie(uid)

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

func (store *EzSqlxStore) GenerateCookieVerifier() (string, string, error) {
	cBytes := make([]byte, 16)
	originalBytes := make([]byte, len(cBytes))

	_, err := rand.Read(cBytes)
	if err != nil {
		return "", "", err
	}

	copy(originalBytes, cBytes)

	hash := sha256.New()
	hash.Write(cBytes)

	return hex.EncodeToString(originalBytes), hex.EncodeToString(hash.Sum(nil)), nil
}

func (store *EzSqlxStore) GenerateCookie(id uuid.UUID) (*ezsesh.EZCookie, string) {
	options := store.Options

	original, verifier, err := store.GenerateCookieVerifier()
	if err != nil {
		return nil, ""
	}

	cookie := &http.Cookie{
		Name:    options.CookieName,
		Value:   hex.EncodeToString([]byte(id.String())) + "-" + verifier,
		Expires: time.Now().Add(time.Duration(options.Lifetime) * time.Minute),

		HttpOnly: options.HttpOnly,
		Secure:   options.Secure,
		SameSite: options.SameSite,
	}

	return &ezsesh.EZCookie{
		Cookie: cookie,
		ID:     id,
	}, original
}

func (store *EzSqlxStore) GetUserSession(assocValue string) (session interface{}, err error) {
	query := fmt.Sprintf("select * from %s where %s = $1", store.Options.Table, store.Options.Association)
	err = store.db.Select(session, query, assocValue)

	return session, err
}

func (store *EzSqlxStore) GetSessionByID(sessionId string) error {
	return nil
}

func (store *EzSqlxStore) DeleteSession(sessionId string) error {
	return nil
}
func (store *EzSqlxStore) DeleteSessionByAssoc(assoc string) error {
	return nil
}
