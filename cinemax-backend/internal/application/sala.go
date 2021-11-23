package application

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (s *service) GetSalas(ctx context.Context) ([]domain.Sala, error) {
	return s.r.GetSalas(ctx)
}
