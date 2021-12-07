package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/middlewares"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util/token"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

type handler struct {
	s ports.Service
	t token.Maker
}

func NewHandler(r *gin.Engine, service ports.Service, token token.Maker) {
	h := &handler{
		s: service,
		t: token,
	}

	// Se hace el registro de las rutas
	h.registerRouter(r)
}

func (h *handler) registerRouter(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		h.s.Ping()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	empleado := e.Group("/empleado")
	empleado.POST("/create", h.CreateEmpleado)
	empleado.POST("/login", h.LoginEmpleado)

	pelicula := e.Group("/pelicula")
	pelicula.Use(middlewares.Auth(h.t))
	pelicula.POST("/registrar", h.CreatePelicula)
	pelicula.POST("/load_image", h.LoadImagen)
	pelicula.POST("/search/id", h.GetPeliculaIdByName)
	pelicula.POST("/search", h.SearchPeliculasByName)
	pelicula.GET("/clasificaciones", h.GetClasificaciones)
	pelicula.GET("/idiomas", h.GetIdiomas)
	pelicula.GET("/generos", h.GetGeneros)
	pelicula.GET("/cartelera", h.GetPeliculasEnCartelera)

	e.GET("/pelicula/image", h.DownloadImagen)

	sala := e.Group("/salas")
	sala.Use(middlewares.Auth(h.t))
	sala.POST("", h.GetSalas)
	sala.GET("/by_funcion", h.GetSalaByFuncionID)

	funcion := e.Group("/funcion")
	funcion.Use(middlewares.Auth(h.t))
	funcion.POST("/create", h.CreateFunction)
	funcion.POST("/show", h.SearchFuncionByPeliculaAndFecha)

	asientos := e.Group("/asiento")
	asientos.Use(middlewares.Auth(h.t))
	asientos.GET("", h.GetAsientosByFuncion)
	asientos.POST("", h.SeleccionarAsiento)
	asientos.DELETE("", h.DeseleccionarAsiento)
	asientos.POST("/deshacer", h.DeshacerTransaccion)

	boleto := e.Group("/boleto")
	boleto.Use(middlewares.Auth(h.t))
	boleto.GET("", h.GetPrecios)

	ticket := e.Group("/ticket")
	ticket.Use(middlewares.Auth(h.t))
	ticket.POST("/iniciar_compra", h.IniciarCompra)
	ticket.POST("/crear_ticket", h.CreateTicket)

}
