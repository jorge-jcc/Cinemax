package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

// Verifica la creación de un token
func TestJWTMarker(t *testing.T) {
	maker := NewJWTMaker("PruebaSecreta")

	id := "1"
	email := "test@test.com"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	// Al crear un token
	token, err := maker.CreateToken(id, email, duration)
	// err debe ser nil
	require.NoError(t, err)
	// token no debe ser un objeto vacío
	require.NotEmpty(t, token)

	// Al validar el token
	payload, err := maker.VerifyToken(token)
	// err debe ser nil
	require.NoError(t, err)
	// payload no debe estar vacío
	require.NotEmpty(t, payload)

	// payload.ID no debe ser cero
	require.NotZero(t, payload.ID)
	// payload.Email debe ser igual a la entrada
	require.Equal(t, email, payload.Email)
	// issuedAt no debe diferir en más de un segundo a payload.IssuedAt
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	// expiredAt no debe diferir en más de un segundo a payload.ExpiredAt
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

// TestExpiredJWTToken verifica el caso de un token expirado
func TestExpiredJWTToken(t *testing.T) {
	maker := NewJWTMaker("PruebaSecreta")

	// Al crear un token
	token, err := maker.CreateToken("1", "email@email.com", -time.Minute)
	// err debe ser nil
	require.NoError(t, err)
	// token no debe ser un objeto vacío
	require.NotEmpty(t, token)

	// Al validar el token
	payload, err := maker.VerifyToken(token)
	// debe obtenerse un error
	require.Error(t, err)
	// el error debe ser ErrExpiredToken
	require.EqualError(t, err, ErrExpiredToken.Error())
	// el payload debe ser nil
	require.Nil(t, payload)
}

// TestInvalidJWTTokenAlgNone verfica el caso de un token invalido
func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// Al crear un nuevo payload
	payload, err := NewPayload("1", "test@test.com", time.Minute)
	// err debe ser nil
	require.NoError(t, err)

	// Al crear un nuevo token con un método de firma diferente
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	// err debe ser nil
	require.NoError(t, err)

	maker := NewJWTMaker("pruebaSecreta")
	// Al verificar el token creado
	payload, err = maker.VerifyToken(token)
	// se debe obtener un error
	require.Error(t, err)
	// el error debe ser ErrInvalidToken
	require.EqualError(t, err, ErrInvalidToken.Error())
	// el payaload debe ser nil
	require.Nil(t, payload)
}
