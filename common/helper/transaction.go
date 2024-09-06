package helper

import (
	"context"
	"database/sql"
)

type TransactionHelper struct {
	db *sql.DB
}

func NewTransactionHelper(db *sql.DB) *TransactionHelper {
	return &TransactionHelper{
		db: db,
	}
}

func (helper *TransactionHelper) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := helper.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = fn(tx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			return errCommit
		}
	}
	return nil
}
