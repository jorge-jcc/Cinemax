package mocks

import (
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util/token"
	"github.com/stretchr/testify/mock"
)

type mockToken struct {
	mock.Mock
}

func NewMockToken() *mockToken {
	return &mockToken{}
}

func (m *mockToken) CreateToken(email string, duration time.Duration) (string, error) {
	args := m.Called(email, duration)    // Lo que debe recibir
	return args.String(0), args.Error(1) // Lo que debe regresar
}

func (m *mockToken) VerifyToken(t string) (*token.Payload, error) {
	args := m.Called(t) // Lo que debe recibir

	var r0 *token.Payload
	if args.Get(0) != nil {
		r0 = args.Get(0).(*token.Payload)
	}
	return r0, args.Error(1)
}
