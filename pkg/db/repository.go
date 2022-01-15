package db

import (
	"errors"

	"github.com/dgraph-io/badger"
)

type BlockchainDB interface {
	Save(key, value []byte, tx *Tx) error
	Find(key []byte) ([]byte, error)
}

// Save stores a new register in the database. If the operation pertences to a transaction
// then set tx with the transaction reference. If doesn't, set tx as null (nil).
func (db *BadgerDB) Save(key, value []byte, tx *Tx) error {

	if tx == nil {
		err := db.Conn.Update(func(txn *badger.Txn) error {

			err := txn.Set(key, value)
			return err

		})

		return err
	}

	err := tx.dbTransaction.Set(key, value)
	return err

}

// Find finds a register in the storage by a given key. If a register is found,
// it will be returned, otherwise, will be returned null (nil).
// If database returns an error, it will be returned by the function.
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
