package ports

import (
	"context"
	"mime/multipart"
	"os"
	"time"

	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type Service interface {
	Ping()

	CreateEmpleado(ctx context.Context, e *domain.Empleado) error
	LoginEmpleado(ctx context.Context, email, password string) (*domain.Empleado, error)

	CreatePelicula(ctx context.Context, p *domain.Pelicula) error
	LoadImage(ctx context.Context, peliculaId string, file *multipart.FileHeader) error
	DownloadImage(ctx context.Context, peliculaId string) (*os.File, error)
	GetPeliculasByNombre(ctx context.Context, nombre string, limit, offset int16) ([]domain.Pelicula, error)
	GetPeliculasEnCartelera(ctx context.Context) ([]domain.Cartelera, error)
	GetClasificaciones(ctx context.Context) ([]domain.Clasificacion, error)
	GetIdiomas(ctx context.Context) ([]domain.Idioma, error)
	GetGeneros(ctx context.Context) ([]domain.Genero, error)

	GetSalas(ctx context.Context, fechaInicio time.Time, peliculaId string) ([]domain.Sala, error)
	GetSalaByFuncionId(ctx context.Context, funcionId string) (*domain.Sala, error)

	CreateFunction(ctx context.Context, horaInicio time.Time, peliculaId, salaId, tipoFuncionId string) error
	GetFuncionesByPeliculaAndFechaInicio(ctx context.Context, peliculaId string, fecha time.Time) ([]domain.Funcion, error)

	GetAsientosByFuncion(ctx context.Context, funcionId string) ([]domain.AsignacionAsiento, error)
	SeleccionarAsiento(ctx context.Context, asientoId string, transaccionId *string) error
	DeseleccionarAsiento(ctx context.Context, asientoId, transaccionId string) error
	DeshacerTransaccion(ctx context.Context, transaccionId string) error

	GetPreciosBoletos(ctx context.Context) ([]domain.PrecioBoleto, error)

	IniciarCompra(ctx context.Context, transaccionId string, a []string) error
	CreateTicket(ctx context.Context, ticket *domain.Ticket) error
}
