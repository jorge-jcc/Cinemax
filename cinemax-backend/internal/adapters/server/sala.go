package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

func (h *handler) GetSalas(c *gin.Context) {
	ctx := c.Request.Context()
	salas, err := h.s.GetSalas(ctx)
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
