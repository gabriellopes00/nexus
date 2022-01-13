package db

import (
	"errors"

	"github.com/dgraph-io/badger"
)

type BlockchainDB interface {
	Save(key, value []byte) error
	Find(key []byte) ([]byte, error)
	Transact(ops ...func(args ...interface{}) error) error
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

// func (db *BadgerDB) Transact(ops ...func(key, value []byte) error) error {
// 	tx := db.Conn.NewTransaction(true)
// 	defer tx.Discard()

// 	for _, op := range ops {
// 		err := op
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err := tx.Commit()
// 	return err
// }

// func (db *BadgerDB) Transact(fns ...func() error) error {

// 	txn := db.Conn.NewTransaction(true)
// 	defer txn.Discard()

// 	for _, fn := range fns {
// 		if err := fn(); err != nil {
// 			return err
// 		}
// 	}

// 	return txn.Commit()
// }
