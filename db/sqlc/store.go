package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v, rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxResult struct {
	Transfers   Transfers `json:"transfer"`
	FromAccount Accounts  `json:"fromAccount"`
	ToAccount   Accounts  `json:"toAccount"`
	FromEntry   Entries   `json:"fromEntry"`
	ToEntry     Entries   `json:"toEntry"`
}

func (store *Store) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result = TransferTxResult{}
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfers, err = q.CreateTransfer(ctx, CreateTransferParams{FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount:      arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.FromAccountID, Amount: -arg.Amount})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.ToAccountID, Amount: arg.Amount})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
