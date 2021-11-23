package application

import (
	"context"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

func (s *service) CreateFunction(ctx context.Context, fechaInicio time.Time, peliculaId, salaId string) error {
	sala, err := s.r.GetSalaById(ctx, salaId)
	if err != nil {
		return err
	}

	p, err := s.r.GetPeliculaById(ctx, peliculaId)
	if err != nil {
		return err
	}

	fechaFin := fechaInicio.Add(time.Minute * time.Duration(p.DuracionMinutos))
	f := &domain.Funcion{
		FechaInicio: fechaInicio,
		FechaFin:    fechaFin,
		Pelicula:    *p,
		Sala:        *sala,
	}

	if ok := s.r.DisponibilidadFuncion(ctx, f); !ok {
		return domain.NewConflict("funcion", salaId)
	}
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		err := s.r.CreateFunction(ctx, f)
		if err != nil {
			return err
		}
		return s.r.InicializarAsientos(c, f)
	})
}

func (s *service) GetFuncionesByPeliculaAndFechaInicio(
	ctx context.Context, peliculaId string, fecha time.Time,
) ([]domain.Funcion, error) {
	return s.r.GetFuncionesByPeliculaAndFechaInicio(ctx, peliculaId, fecha)
}
