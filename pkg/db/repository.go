package db

import (
	"errors"

	"github.com/dgraph-io/badger"
)

type BlockchainDB interface {
	Save(key, value []byte) error
	Find(key []byte) (value []byte, err error)
}

func (db *BadgerDB) Save(key, value []byte) error {
	err := db.Conn.Update(func(txn *badger.Txn) error {

		err := txn.Set(key, value)
		return err

	})

	return err
}

func (db *BadgerDB) Find(key []byte) (value []byte, err error) {
	err = db.Conn.View(func(txn *badger.Txn) error {

		item, err := txn.Get(key)
		if err != nil {

			if errors.Is(err, badger.ErrKeyNotFound) {
				value = nil
				return nil
			}

			return err
		}

		value, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	return
}
