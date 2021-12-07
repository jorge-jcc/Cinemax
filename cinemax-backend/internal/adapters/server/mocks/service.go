package mocks

import (
	"context"
	"mime/multipart"
	"os"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func NewMockService() *mockService {
	return &mockService{}
}

func (m *mockService) CreateEmpleado(ctx context.Context, e *domain.Empleado) error {
	args := m.Called(ctx, e) // Lo que debe recibir
	return args.Error(0)     // Lo que debe regresar
}

func (m *mockService) LoginEmpleado(ctx context.Context, email, password string) (*domain.Empleado, error) {
	args := m.Called(ctx, email, password) // Lo que debe recibir

	var r0 *domain.Empleado
	if args.Get(0) != nil {
		r0 = args.Get(0).(*domain.Empleado)
	}

	var r1 error
	if args.Get(1) != nil {
		r1 = args.Error(1)
	}
	return r0, r1 // Lo que debe regresar
}

func (m *mockService) Ping() {}

func (m *mockService) CreatePelicula(ctx context.Context, p *domain.Pelicula) error {
	return nil
}

func (m *mockService) LoadImage(ctx context.Context, peliculaId string, file *multipart.FileHeader) error {
	return nil
}

func (s *mockService) DownloadImage(ctx context.Context, peliculaId string) (*os.File, error) {
	return nil, nil
}

func (s *mockService) GetPeliculasByNombre(ctx context.Context, nombre string, offset, limit int16) ([]domain.Pelicula, error) {
	return nil, nil
}

func (s *mockService) GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error) {
	return nil, nil
}

func (s *mockService) GetIdiomas(ctx context.Context) ([]domain.Idioma, error) {
	return nil, nil
}

func (s *mockService) GetGeneros(ctx context.Context) ([]domain.Genero, error) {
	return nil, nil
}

func (s *mockService) GetSalas(ctx context.Context) ([]domain.Sala, error) {
	return nil, nil
}
func (s *mockService) CreateFunction(ctx context.Context, horaInicio time.Time, peliculaId, salaId, tipoFuncionId string) error {
	return nil
}
func (s *mockService) GetFuncionesByPeliculaAndFechaInicio(ctx context.Context, peliculaId string, fecha time.Time) ([]domain.Funcion, error) {
	return nil, nil
}
func (s *mockService) GetAsientosByFuncion(ctx context.Context, funcionId string) ([]domain.AsignacionAsiento, error) {
	return nil, nil
}
func (s *mockService) SeleccionarAsiento(ctx context.Context, asientoId string, transaccionId *string) error {
	return nil
}

func (s *mockService) GetPeliculasEnCartelera(ctx context.Context) ([]domain.Pelicula, error) {
	return nil, nil
}

func (s *mockService) GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error) {
	return nil, nil
}

func (s *mockService) DeseleccionarAsiento(ctx context.Context, asientoId, transaccionId string) error {
	return nil
}

func (s *mockService) DeshacerTransaccion(ctx context.Context, transaccionId string) error {
	return nil
}

func (s *mockService) GetPreciosBoletos(ctx context.Context) ([]domain.PrecioBoleto, error) {
	return nil, nil
}
