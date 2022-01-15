package db

import "github.com/dgraph-io/badger"

type Tx struct {
	// dbTransaction maps to the badger transaction struct.
	dbTransaction *badger.Txn
}

func newTx(dbTransaction *badger.Txn) *Tx {
	return &Tx{
		dbTransaction: dbTransaction,
	}
}

// Commit commits the changes in the storage. It returns an error if there is one.
func (tx *Tx) Commit() error {
	return tx.dbTransaction.Commit()
}

// Rollback cancel all the operations of the transaction, discarding any change
// made in the storage. Once rolledback, a transaction cannot be used again.
func (tx *Tx) Rollback() {
	tx.dbTransaction.Discard()
}
