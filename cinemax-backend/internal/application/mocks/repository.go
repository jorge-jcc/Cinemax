package mocks

import (
	"context"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func NewMockRepository() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) Ping() error {
	return nil
}

func (m *mockRepository) Transaction(ctx context.Context, tx ports.TxFn) error {
	args := m.Called(ctx, tx) // Lo que debe recibir
	return args.Error(0)      // Lo que debe regresar
}

func (m *mockRepository) CreateEmpleado(ctx context.Context, e *domain.Empleado) error {
	m.Called(ctx, e)
	return nil
}

func (m *mockRepository) FindEmpleadoByEmail(ctx context.Context, email string) (*domain.Empleado, error) {
	args := m.Called(ctx, email)

	var r0 *domain.Empleado
	if args.Get(0) != nil {
		r0 = args.Get(0).(*domain.Empleado)
	}
	return r0, args.Error(1)
}
func (m *mockRepository) CreatePelicula(ctx context.Context, p *domain.Pelicula) error {
	return nil
}

func (m *mockRepository) UpdateImage(ctx context.Context, id, imagen string) error {
	return nil
}

func (r *mockRepository) GetPeliculaById(ctx context.Context, id string) (*domain.Pelicula, error) {
	return nil, nil
}

func (r *mockRepository) GetPeliculasByNombre(ctx context.Context, nombre string, offset, limit int16) ([]domain.Pelicula, error) {
	return nil, nil
}

func (r *mockRepository) GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error) {
	return nil, nil
}

func (s *mockRepository) GetIdiomas(ctx context.Context) ([]domain.Idioma, error) {
	return nil, nil
}

func (s *mockRepository) GetGeneros(ctx context.Context) ([]domain.Genero, error) {
	return nil, nil
}

func (s *mockRepository) GetSalas(ctx context.Context) ([]domain.Sala, error) {
	return nil, nil
}

func (s *mockRepository) CreateFunction(ctx context.Context, f *domain.Funcion) error {
	return nil
}

func (s *mockRepository) GetSalaById(ctx context.Context, id string) (*domain.Sala, error) {
	return nil, nil
}

func (s *mockRepository) DisponibilidadFuncion(ctx context.Context, f *domain.Funcion) bool {
	return false
}

func (s *mockRepository) GetFuncionesByPeliculaAndFechaInicio(ctx context.Context, peliculaId string, fecha time.Time) ([]domain.Funcion, error) {
	return nil, nil
}
