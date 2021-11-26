package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

type querier interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type repository struct {
	db querier
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r repository) Transaction(ctx context.Context, txFn ports.TxFn) error {
	tx, _ := r.db.(*sqlx.DB).Beginx()
	r.db = tx
	err := txFn(ctx, &r)
	if err != nil {
		if tx.Rollback() != nil {
			return fmt.Errorf("failed to execute transaction, %w", err)
		}
		return err
	}
	if tx.Commit() != nil {
		return fmt.Errorf("failed to commit transaction, %w", err)
	}
	return nil
}

func (r *repository) Ping() error {
	if db, ok := r.db.(*sqlx.DB); ok {
		return db.Ping()
	}
	return fmt.Errorf("operation not supported")
}
