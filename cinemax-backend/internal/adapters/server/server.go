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
	pelicula.GET("/image", h.DownloadImagen)
	pelicula.POST("/search", h.SearchPeliculasByName)
	pelicula.GET("/clasificaciones", h.GetClasificaciones)
	pelicula.GET("/idiomas", h.GetIdiomas)
	pelicula.GET("/generos", h.GetGeneros)

	sala := e.Group("/sala")
	sala.Use(middlewares.Auth(h.t))
	sala.GET("", h.GetSalas)

	funcion := e.Group("/funcion")
	funcion.Use(middlewares.Auth(h.t))
	funcion.POST("/create", h.CreateFunction)
	funcion.POST("/show", h.SearchFuncionByPeliculaAndFecha)

	asientos := e.Group("/asiento")
	asientos.Use(middlewares.Auth(h.t))
	asientos.GET("", h.GetAsientosByFuncion)
	asientos.POST("/seleccionar", h.SeleccionarAsiento)

	taquilla := e.Group("/taquilla")
	taquilla.Use(middlewares.Auth(h.t))
	taquilla.GET("/ping", func(c *gin.Context) {
		h.s.Ping()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
