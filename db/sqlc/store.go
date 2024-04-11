package db

import (
	"database/sql"
)

// embedding (compostition + interface instead of inheritance)
type IStore interface {
	Querier
	//TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, globalErr error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

//var TkKey = struct{}{}

func NewStore(db *sql.DB) IStore {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

//// execute a function within a database transaction
//func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
//	tx, err := s.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	q := New(tx)
//	err = fn(q)
//	if err != nil {
//		if rbErr := tx.Rollback(); rbErr != nil {
//			return fmt.Errorf("%v AND %v", err, rbErr)
//		}
//		return err
//	}
//	return tx.Commit()
//}
