package db

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

var db *badger.DB

func Init(path string) {
	var err error
	opts := badger.DefaultOptions(path).
		WithLoggingLevel(badger.WARNING)
	db, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
}
