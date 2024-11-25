package helper

import (
	"database/sql"
	// "fmt"
)

type transaction struct {
	DB *sql.DB
}

type TransactionFunc interface {
	BeginTransaction() (*sql.Tx, error)
	RollbackTransaction(tx *sql.Tx) (error)
	CommitTransaction(tx *sql.Tx) (error)	
}

func CreateTransaction(db *sql.DB) TransactionFunc {
    return &transaction{DB: db}
}

func (trx *transaction) BeginTransaction() (*sql.Tx, error) {
	tx, err := trx.DB.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (trx *transaction) RollbackTransaction(tx *sql.Tx) (error) {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (trx *transaction) CommitTransaction(tx *sql.Tx) (error) {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
