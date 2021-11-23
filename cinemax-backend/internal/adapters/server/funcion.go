package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type createFuncionReq struct {
	HoraInicio time.Time `json:"horaInicio" binding:"required"`
	PeliculaId string    `json:"peliculaId" binding:"required"`
	SalaID     string    `json:"salaId" binding:"required"`
}

type searchFuncionReq struct {
	PeliculaID string    `json:"peliculaId"`
	Fecha      time.Time `json:"fecha"`
}

func (h *handler) CreateFunction(c *gin.Context) {
	// Enlaza el JSON entrante a estructura req y verificar errores de validación
	var req createFuncionReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	err := h.s.CreateFunction(ctx, req.HoraInicio, req.PeliculaId, req.SalaID)
	if err != nil {
		log.Printf("Failed to create funcion: %v\n", err.Error())
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *handler) SearchFuncionByPeliculaAndFecha(c *gin.Context) {
	var req searchFuncionReq
	if ok := util.BindData(c, &req); !ok {
		return
	}
	if req.Fecha.IsZero() {
		req.Fecha = time.Now()
	}

	ctx := c.Request.Context()

	funciones, err := h.s.GetFuncionesByPeliculaAndFechaInicio(ctx, req.PeliculaID, req.Fecha)
	if err != nil {
		log.Printf("Failed to get funciones: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"funciones": funciones,
	})
}
