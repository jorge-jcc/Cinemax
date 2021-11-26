package ports

import (
	"context"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type TxFn func(context.Context, Repository) error

type Repository interface {
	Ping() error
	Transaction(ctx context.Context, tx TxFn) error

	CreateEmpleado(ctx context.Context, e *domain.Empleado) error
	FindEmpleadoByEmail(ctx context.Context, email string) (*domain.Empleado, error)

	CreatePelicula(ctx context.Context, p *domain.Pelicula) error
	UpdateImage(ctx context.Context, id, imagen string) error
	GetPeliculaById(ctx context.Context, id string) (*domain.Pelicula, error)
	GetPeliculasByNombre(ctx context.Context, nombre string, limit, offset int16) ([]domain.Pelicula, error)
	GetPeliculasEnCartelera(ctx context.Context) ([]domain.Pelicula, error)
	GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error)
	GetIdiomas(ctx context.Context) ([]domain.Idioma, error)
	GetGeneros(ctx context.Context) ([]domain.Genero, error)

	GetSalas(ctx context.Context) ([]domain.Sala, error)
	GetSalaById(ctx context.Context, id string) (*domain.Sala, error)
	GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error)

	CreateFunction(ctx context.Context, f *domain.Funcion) error
	DisponibilidadFuncion(ctx context.Context, f *domain.Funcion) bool
	GetFuncionesByPeliculaAndFechaInicio(ctx context.Context, peliculaId string, fecha time.Time) ([]domain.Funcion, error)

	InicializarAsientos(ctx context.Context, f *domain.Funcion) error
	GetAsientosByFuncion(ctx context.Context, funcionId string) ([]domain.AsignacionAsiento, error)
	GetAsientoByID(ctx context.Context, asientoId string) (*domain.AsignacionAsiento, error)
	DisponibilidadAsiento(ctx context.Context, a *domain.AsignacionAsiento) bool
	UpdateStatusAsiento(ctx context.Context, a *domain.AsignacionAsiento) error

	GetNewTransactionID(ctx context.Context) (string, error)
	ValidarTransaccion(ctx context.Context, transaccionId string) error
	UpdateTimeTransaction(ctx context.Context, transaccionId string) error
	DeshacerTransaccion(ctx context.Context, transaccionId string) error
}
