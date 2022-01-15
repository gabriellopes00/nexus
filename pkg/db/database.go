package db

import (
	"nexus/env"
	"os"

	"github.com/dgraph-io/badger"
)

var dbPath = env.DATABASE_PATH + env.DATABASE_FILE

type BadgerDB struct {
	Conn *badger.DB
}

// NewBadgerDB returns a new instance of the badger database.
func NewBadgerDB() (*BadgerDB, error) {
	wd, _ := os.Getwd()

	options := badger.DefaultOptions(wd + dbPath)
	options.ValueDir = dbPath
	options.Logger = nil

	connection, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	return &BadgerDB{Conn: connection}, nil
}

// NewTransaction returns a new transaction. For read-write operations, set write
// equal to true.
func (db *BadgerDB) NewTransaction(write bool) *Tx {
	dbTransaction := db.Conn.NewTransaction(write)
	return newTx(dbTransaction)
}
