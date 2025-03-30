package provider

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lecterkn/goat_backend/internal/app/database"
	"github.com/lecterkn/goat_backend/internal/app/port"
)

type TransactionProviderImpl struct {
	database *sqlx.DB
}

func NewTransactionProviderImpl(database *sqlx.DB) port.TransactionProvider {
	return &TransactionProviderImpl{
		database,
	}
}

func (p *TransactionProviderImpl) Transact(txFn func(ctx context.Context) error) error {
	ctx := context.Background()
	tx, err := p.database.Beginx()
	if err != nil {
		return err
	}
	p.injectTx(ctx, tx)
	err = txFn(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (p *TransactionProviderImpl) injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, database.TxKey, tx)
}
