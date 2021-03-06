package application

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

func (s *service) CreatePelicula(ctx context.Context, p *domain.Pelicula) error {
	err := s.r.Transaction(ctx,
		func(ctx context.Context, tx ports.Repository) error {
			return tx.CreatePelicula(ctx, p)
		})
	return err
}

func (s *service) LoadImage(ctx context.Context, id string, file *multipart.FileHeader) error {
	extension := strings.Split(file.Filename, ".")[1]
	fileName := fmt.Sprintf("pelicula_%s.%s", id, extension)
	return s.r.Transaction(ctx, func(c context.Context, r ports.Repository) error {
		err := r.UpdateImage(c, id, fileName)
		if err != nil {
			return err
		}
		return s.i.LoadImage(c, fileName, file)
	})
}

func (s *service) DownloadImage(ctx context.Context, peliculaId string) (*os.File, error) {
	p, err := s.r.GetPeliculaById(ctx, peliculaId)
	if err != nil {
		return nil, err
	}

	file, err := s.i.DownloadImage(ctx, p.Imagen)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *service) GetPeliculasByNombre(ctx context.Context, nombre string, limit, offset int16) ([]domain.Pelicula, error) {
	return s.r.GetPeliculasByNombre(ctx, nombre, limit, offset)
}

func (s *service) GetPeliculasEnCartelera(ctx context.Context) ([]domain.Cartelera, error) {
	p, err := s.r.GetPeliculasEnCartelera(ctx)
	if err != nil {
		return nil, err
	}
	c := make([]domain.Cartelera, len(p))
	for i := range p {
		f, err := s.r.GetFuncionesByPeliculaAndFechaInicio(ctx, p[i].ID, time.Now())
		if err != nil {
			return nil, err
		}
		c[i].Pelicula = p[i]
		for j := range f {
			c[i].Horarios = append(c[i].Horarios, f[j].FechaInicio.Format("3:04 PM"))
		}
	}
	return c, nil
}

func (s *service) GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error) {
	return s.r.GetClasificaciones(ctx)
}

func (s *service) GetIdiomas(ctx context.Context) ([]domain.Idioma, error) {
	return s.r.GetIdiomas(ctx)
}

func (s *service) GetGeneros(ctx context.Context) ([]domain.Genero, error) {
	return s.r.GetGeneros(ctx)
}
