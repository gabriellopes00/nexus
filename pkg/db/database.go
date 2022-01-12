package db

import (
	"os"

	"github.com/dgraph-io/badger"
)

const dbPath = ".\\temp\\db"

type BadgerDB struct {
	Conn *badger.DB
}

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
