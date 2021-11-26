package application

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (s *service) GetSalas(ctx context.Context) ([]domain.Sala, error) {
	return s.r.GetSalas(ctx)
}

func (s *service) GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error) {
	return s.r.GetSalaByFuncionId(ctx, funcionId)
}
