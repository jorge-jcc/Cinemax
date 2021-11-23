package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/mocks"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/stretchr/testify/mock"
)

func TestCreateEmpleado(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()

	mockService := mocks.NewMockService()
	mockToken := mocks.NewMockToken()
	NewHandler(router, mockService, mockToken)

	var mockArgs mock.Arguments
	var mockError *domain.Error

	// Usuario no existente
	mockArgs = mock.Arguments{
		mock.AnythingOfType("*context.emptyCtx"),
		"notuser@test.com", "password",
	}
	mockError = domain.NewAuthorizationError("invalid email/password combo")
	mockService.On("LoginEmpleado", mockArgs...).Return(nil, mockError)

	// Password incorrecto
	mockArgs = mock.Arguments{
		mock.AnythingOfType("*context.emptyCtx"),
		"user@test.com", "notpassword",
	}
	mockError = domain.NewAuthorizationError("invalid email/password combo")
	mockService.On("LoginEmpleado", mockArgs...).Return(nil, mockError)

	// Usuario y Contraseña correctos
	// Error al generar token
	mockArgs = mock.Arguments{
		mock.AnythingOfType("*context.emptyCtx"),
		"nottoken@test.com", "password",
	}
	mockService.On("LoginEmpleado", mockArgs...).Return(&domain.Empleado{Email: "nottoken@test.com"}, nil)
	mockToken.On("CreateToken", "nottoken@test.com", time.Hour*24).Return("", mockError)
	// Token generado correctamente
	mockArgs = mock.Arguments{
		mock.AnythingOfType("*context.emptyCtx"),
		"token@test.com", "password",
	}
	mockService.On("LoginEmpleado", mockArgs...).Return(&domain.Empleado{Email: "token@test.com"}, nil)
	mockToken.On("CreateToken", "token@test.com", time.Hour*24).Return("token", nil)

	testCases := []util.TestVector{
		{
			Name: "BadRequestData",
			Internal: []util.TestVector{
				{
					Name: "EmailInvalido",
					Body: gin.H{
						"email":    "email_invalido",
						"password": "short",
					},
					StatusCode: http.StatusBadRequest,
				},
				{
					Name: "PasswordInvalido",
					Body: gin.H{
						"email":    "test@test.com",
						"password": "short",
					},
					StatusCode: http.StatusBadRequest,
				},
			},
		},
		{
			Name: "UsuarioOContraseñaIncorrectos",
			Internal: []util.TestVector{
				{
					Name: "UsuarioNotFound",
					Body: gin.H{
						"email":    "notuser@test.com",
						"password": "password",
					},
					StatusCode: http.StatusUnauthorized,
				},
				{
					Name: "ContraseñaIncorrecto",
					Body: gin.H{
						"email":    "user@test.com",
						"password": "notpassword",
					},
					StatusCode: http.StatusUnauthorized,
				},
			},
		},
		{
			Name: "UsuarioYContraseñaCorrectos",
			Internal: []util.TestVector{
				{
					Name: "ErrorAlGenerarToken",
					Body: gin.H{
						"email":    "nottoken@test.com",
						"password": "password",
					},
					StatusCode: http.StatusInternalServerError,
				},
				{
					Name: "TokenExitoso",
					Body: gin.H{
						"email":    "token@test.com",
						"password": "password",
					},
					StatusCode: http.StatusOK,
				},
			},
		},
	}

	util.RunTest(t, router, "POST", "/empleado/login", testCases)
}
