package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type salasReq struct {
	PeliculaId  string    `json:"peliculaId" binding:"required"`
	FechaInicio time.Time `json:"fechaInicio" binding:"required"`
}

func (h *handler) GetSalas(c *gin.Context) {
	var req salasReq
	if ok := util.BindData(c, &req); !ok {
		return
	}
	ctx := c.Request.Context()
	salas, err := h.s.GetSalas(ctx, req.FechaInicio, req.PeliculaId)
	if err != nil {
		log.Printf("Failed to get salas %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"salas": salas,
	})
}

func (h *handler) GetSalaByFuncionID(c *gin.Context) {
	var funcionId string
	if ok := util.BindQuery(c, "funcion_id", &funcionId); !ok {
		return
	}
	ctx := c.Request.Context()
	sala, err := h.s.GetSalaByFuncionId(ctx, funcionId)
	if err != nil {
		log.Printf("Failed to get salas %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sala": sala,
	})
}
