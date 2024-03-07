package ezsesh

import "github.com/jmoiron/sqlx"

type EzStore struct {
	EzStoreMethods
	options *EzOptions
	db      *sqlx.DB
}

func (store *EzStore) Generate() {

}
