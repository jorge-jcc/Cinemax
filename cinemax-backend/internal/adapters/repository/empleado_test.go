package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestFindEmpleadoByEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	repository := NewRepository(sqlx.NewDb(db, "postgres"))

	ctx := context.TODO()

	// Caso 1: el usuario no existe
	t.Run("EmailNoRegistrado", func(t *testing.T) {
		mock.ExpectQuery(`SELECT * FROM "EMPLEADO" WHERE "EMAIL" = $1`).
			WithArgs("notuser@test.com").WillReturnError(sql.ErrNoRows)
		e, err := repository.FindEmpleadoByEmail(ctx, "notuser@test.com")
		require.Error(t, err)
		require.EqualError(t, err, domain.NewNotFound("email", "notuser@test.com").Error())
		require.Nil(t, e)
	})
	// Caso 2: el usuario existe
	t.Run("UsuarioEncontrado", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"EMPLEADO_ID"}).AddRow("1")
		mock.ExpectQuery(`SELECT * FROM "EMPLEADO" WHERE "EMAIL" = $1`).
			WithArgs("user@test.com").WillReturnRows(rows)
		e, err := repository.FindEmpleadoByEmail(ctx, "user@test.com")
		require.Nil(t, err)
		require.NotNil(t, e)
	})
}
