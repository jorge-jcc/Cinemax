package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (h *handler) GetPrecios(c *gin.Context) {
	ctx := c.Request.Context()
	salas, err := h.s.GetPreciosBoletos(ctx)
	if err != nil {
		log.Printf("Failed to get precios %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"precios": salas,
	})
}
