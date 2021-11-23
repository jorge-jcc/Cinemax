package application

import (
	"context"
	"log"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/application/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

func (s *service) CreateEmpleado(ctx context.Context, e *domain.Empleado) error {
	pw, err := util.HashPassword(e.Password)
	if err != nil {
		log.Printf("Unable to signup user for email: %v\n", e.Email)
		return domain.NewInternal()
	}
	e.Password = pw
	return s.r.Transaction(ctx,
		func(ctx context.Context, tx ports.Repository) error {
			return tx.CreateEmpleado(ctx, e)
		})
}

func (s *service) LoginEmpleado(ctx context.Context, email, password string) (*domain.Empleado, error) {
	e, err := s.r.FindEmpleadoByEmail(ctx, email)
	if err != nil {
		return nil, domain.NewAuthorizationError("El email o contraseña no son validos")
	}
	ok := util.AuthenticatePassword(e.Password, password)
	if !ok {
		return nil, domain.NewAuthorizationError("El email o contraseña no son validos")
	}
	return e, nil
}
