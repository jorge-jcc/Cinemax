package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type seleccionarAsientoReq struct {
	AsientoId     string `json:"asientoId" binding:"required"`
	TransaccionId string `json:"transaccionId"`
}

func (h *handler) GetAsientosByFuncion(c *gin.Context) {
	var reqFuncionId string
	if ok := util.BindQuery(c, "funcion_id", &reqFuncionId); !ok {
		return
	}

	ctx := c.Request.Context()

	asientos, err := h.s.GetAsientosByFuncion(ctx, reqFuncionId)
	if err != nil {
		log.Printf("Failed to get funciones: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"asientos": asientos,
	})
}

func (h *handler) SeleccionarAsiento(c *gin.Context) {
	var req seleccionarAsientoReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	err := h.s.SeleccionarAsiento(ctx, req.AsientoId, &req.TransaccionId)
	if err != nil {
		log.Printf("Failed to select Asiento: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"transaccionId": req.TransaccionId,
	})
}

func (h *handler) DeseleccionarAsiento(c *gin.Context) {
	type transaccionIdReq struct {
		TransaccionId string `json:"transaccionId" binding:"required"`
		AsientoID     string `json:"asientoId" binding:"required"`
	}
	var req transaccionIdReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	err := h.s.DeseleccionarAsiento(ctx, req.AsientoID, req.TransaccionId)
	if err != nil {
		log.Printf("Failed to deseleccionar asiento: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *handler) DeshacerTransaccion(c *gin.Context) {
	type transaccionIdReq struct {
		TransaccionId string `json:"transaccionId" binding:"required"`
	}
	var req transaccionIdReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	err := h.s.DeshacerTransaccion(ctx, req.TransaccionId)
	if err != nil {
		log.Printf("Error al deshacer la transaccion: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.Status(http.StatusNoContent)
}
