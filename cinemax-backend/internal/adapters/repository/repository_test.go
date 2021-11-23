package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	t.Run("CommitSuccess", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close()

		repository := NewRepository(sqlx.NewDb(db, "postgres"))
		ctx := context.TODO()
		mock.ExpectBegin()
		mock.ExpectCommit()
		err = repository.Transaction(ctx, func(c context.Context, r ports.Repository) error {
			return nil
		})
		require.NoError(t, err)
		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
	t.Run("RollbackSuccess", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close()

		repository := NewRepository(sqlx.NewDb(db, "postgres"))
		ctx := context.TODO()
		mock.ExpectBegin()
		mock.ExpectRollback()
		err = repository.Transaction(ctx, func(c context.Context, r ports.Repository) error {
			return errors.New("any error")
		})
		require.Error(t, err)
		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
}
