package server

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type reqPelicula struct {
	Nombre             string    `json:"nombre" binding:"required"`
	Director           string    `json:"director" binding:"required"`
	Descripcion        string    `json:"descripcion" binding:"required"`
	DuracionMinutos    int16     `json:"duracionMinutos" binding:"required,numeric"`
	Anio               string    `json:"anio" binding:"required"`
	FechaDisponiblidad time.Time `json:"fechaDisponibilidad" binding:"required"`
	Resena             string    `json:"resena" binding:"required"`
	ClasificacionId    string    `json:"clasificacionId" binding:"required"`
	IdiomaId           string    `json:"idiomaId" binding:"required"`
	SubtitiuloId       string    `json:"subtituloId" binding:"required"`
	GeneroId           string    `json:"generoId" binding:"required"`
}

type reqSearchPelicula struct {
	Nombre string `json:"nombre"`
	Limit  int16  `json:"limit" binding:"numeric"`
	Offset int16  `json:"offset" binding:"numeric"`
}

func (h *handler) CreatePelicula(c *gin.Context) {
	// Validando informacion de la pelicula
	var req reqPelicula
	if ok := util.BindData(c, &req); !ok {
		return
	}

	p := domain.NewPelicula(
		req.Nombre,
		req.Director,
		req.Descripcion,
		req.DuracionMinutos,
		req.Anio,
		req.FechaDisponiblidad,
		req.Resena,
		req.ClasificacionId,
		req.IdiomaId,
		req.SubtitiuloId,
		req.GeneroId,
	)

	ctx := c.Request.Context()

	err := h.s.CreatePelicula(ctx, p)
	if err != nil {
		log.Printf("Failed to create pelicula: %v\n", err.Error())
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"peliculaId": p.ID,
	})
}

func (h *handler) GetPeliculaIdByName(c *gin.Context) {
	type response struct {
		Id     string `json:"id"`
		Nombre string `json:"nombre"`
	}

	var req reqSearchPelicula
	if ok := util.BindData(c, &req); !ok {
		return
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	ctx := c.Request.Context()

	peliculas, err := h.s.GetPeliculasByNombre(ctx, req.Nombre, req.Limit, req.Offset)
	if err != nil {
		log.Printf("Failed to get peliculas: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	p := make([]response, len(peliculas))
	for i := range peliculas {
		p[i].Id = peliculas[i].ID
		p[i].Nombre = peliculas[i].Nombre
	}

	c.JSON(http.StatusOK, gin.H{
		"peliculas": p,
	})
}

func (h *handler) SearchPeliculasByName(c *gin.Context) {
	var req reqSearchPelicula
	if ok := util.BindData(c, &req); !ok {
		return
	}
	if req.Limit == 0 {
		req.Limit = 20
	}

	ctx := c.Request.Context()

	peliculas, err := h.s.GetPeliculasByNombre(ctx, req.Nombre, req.Limit, req.Offset)
	if err != nil {
		log.Printf("Failed to get peliculas: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"peliculas": peliculas,
	})
}

func (h *handler) GetPeliculasEnCartelera(c *gin.Context) {
	type resPelicula struct {
		ID              string   `json:"id"`
		Nombre          string   `json:"nombre"`
		Clasificacion   string   `json:"clasificacion"`
		DuracionMinutos int16    `json:"duracionMinutos"`
		Genero          string   `json:"genero"`
		Idioma          string   `json:"idioma"`
		Subtitulo       string   `json:"subtitulo"`
		Horarios        []string `json:"horarios"`
	}
	ctx := c.Request.Context()

	cartelera, err := h.s.GetPeliculasEnCartelera(ctx)
	if err != nil {
		log.Printf("Failed to get cartelera: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	ca := make([]resPelicula, len(cartelera))
	for i := range cartelera {
		ca[i].ID = cartelera[i].ID
		ca[i].Nombre = cartelera[i].Nombre
		ca[i].Clasificacion = cartelera[i].Clasificacion.Clave
		ca[i].DuracionMinutos = cartelera[i].DuracionMinutos
		ca[i].Genero = cartelera[i].Genero.Nombre
		ca[i].Idioma = cartelera[i].Idioma.Nombre
		ca[i].Subtitulo = cartelera[i].Subtitulo.Nombre
		ca[i].Horarios = cartelera[i].Horarios
	}
	c.JSON(http.StatusOK, gin.H{
		"peliculas": ca,
	})
}

func (h *handler) LoadImagen(c *gin.Context) {
	// Validando el Id de la imagen
	peliculaId := c.PostForm("peliculaId")
	if peliculaId == "" {
		e := domain.NewBadRequest("the movie id was expected")
		c.JSON(domain.Status(e), gin.H{
			"error": e,
		})
		return
	}

	// Validando la imagen
	var reqImage *multipart.FileHeader
	if ok := util.BindFile(c, "portada", &reqImage); !ok {
		return
	}

	ctx := c.Request.Context()

	err := h.s.LoadImage(ctx, peliculaId, reqImage)
	if err != nil {
		log.Printf("Failed to save image file: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *handler) DownloadImagen(c *gin.Context) {
	var req string
	if ok := util.BindQuery(c, "pelicula_id", &req); !ok {
		return
	}

	ctx := c.Request.Context()

	imageFile, err := h.s.DownloadImage(ctx, req)
	if err != nil {
		log.Printf("Failed to download image file: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	defer imageFile.Close()

	var f io.Reader = imageFile
	_, err = io.Copy(c.Writer, f)
	if err != nil {
		e := domain.NewNotFound("imagen", "")
		log.Printf("Failed to download image file: %v\n", err)
		c.JSON(domain.Status(e), gin.H{
			"error": err,
		})
	}
}

func (h *handler) GetClasificaciones(c *gin.Context) {
	ctx := c.Request.Context()
	clasificaciones, err := h.s.GetClasificaciones(ctx)
	if err != nil {
		log.Printf("Failed to get clasificaciones %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"clasificaciones": clasificaciones,
	})
}

func (h *handler) GetIdiomas(c *gin.Context) {
	ctx := c.Request.Context()
	idiomas, err := h.s.GetIdiomas(ctx)
	if err != nil {
		log.Printf("Failed to get idiomas %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"idiomas": idiomas,
	})
}

func (h *handler) GetGeneros(c *gin.Context) {
	ctx := c.Request.Context()
	generos, err := h.s.GetGeneros(ctx)
	if err != nil {
		log.Printf("Failed to get generos %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"generos": generos,
	})
}
