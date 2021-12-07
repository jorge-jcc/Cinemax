package application

import (
	"context"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (s *service) GetPreciosBoletos(ctx context.Context) ([]domain.PrecioBoleto, error) {
	return s.r.GetPreciosBoletos(ctx)
}
