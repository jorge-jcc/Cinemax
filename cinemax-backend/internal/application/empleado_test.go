package application

import (
	"context"
	"errors"
	"testing"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/application/mocks"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLoginEmpleado(t *testing.T) {
	mockRepository := mocks.NewMockRepository()
	service := NewService(mockRepository, nil)

	ctx := context.TODO()

	// Case 1: el usuario no existe en la base de datos
	mockRepository.On("FindEmpleadoByEmail", mock.Anything, "notuser@test.com").
		Return(nil, errors.New("any error"))

	// Case 2: el usuario existe
	e := &domain.Empleado{
		Id:       "1",
		Email:    "user@test.com",
		Password: "$2a$10$zYkZszMQQuKii3byGELt7OWQbOb/9E2YSKXlxEykSrRKUiEd/BJea",
	}
	mockRepository.On("FindEmpleadoByEmail", mock.Anything, "user@test.com").
		Return(e, nil)

	t.Run("ElUsuarioNoExiste", func(t *testing.T) {
		e, err := service.LoginEmpleado(ctx, "notuser@test.com", "password")
		require.Error(t, err)
		require.EqualError(t, err, domain.NewAuthorizationError("El email o contraseña no son validos").Error())
		require.Nil(t, e)
	})
	t.Run("LaContraseñaNoEsCorrecta", func(t *testing.T) {
		e, err := service.LoginEmpleado(ctx, "user@test.com", "password")
		require.Error(t, err)
		require.EqualError(t, err, domain.NewAuthorizationError("El email o contraseña no son validos").Error())
		require.Nil(t, e)
	})
	t.Run("UsuarioYContraseñaCorrectos", func(t *testing.T) {
		e, err := service.LoginEmpleado(ctx, "user@test.com", "123456")
		require.Nil(t, err)
		require.NotNil(t, e)
	})
}
