package application

import (
	"context"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (s *service) GetSalas(ctx context.Context, fechaInicio time.Time, peliculaId string) ([]domain.Sala, error) {
	p, err := s.r.GetPeliculaById(ctx, peliculaId)
	if err != nil {
		return nil, err
	}
	return s.r.GetSalasDisponibles(ctx, fechaInicio, fechaInicio.Add(time.Duration(p.DuracionMinutos)*time.Minute))
}

func (s *service) GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error) {
	return s.r.GetSalaByFuncionId(ctx, funcionId)
}
