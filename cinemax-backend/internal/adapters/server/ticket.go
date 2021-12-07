package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util/token"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type boletoReq struct {
	TipoBoletoID string `json:"tipoBoletoId"`
	AsientoId    string `json:"asientoId"`
}

type ticketReq struct {
	TransaccionId string      `json:"transaccionId" binding:"required"`
	FuncionId     string      `json:"funcionId" binding:"required"`
	Monto         float32     `json:"monto" binding:"required"`
	Boletos       []boletoReq `json:"boletos" binding:"required"`
}

type inicarCompraReq struct {
	TransaccionId string   `json:"transaccionId" binding:"required"`
	Boletos       []string `json:"boletos" binding:"required"`
}

func (h *handler) CreateTicket(c *gin.Context) {
	var req ticketReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	e, _ := c.Get("empleado")
	boletos := make([]domain.Boleto, len(req.Boletos))
	for i := range req.Boletos {
		boletos[i].TipoBoletoId = req.Boletos[i].TipoBoletoID
		boletos[i].AsientoId = req.Boletos[i].AsientoId
	}
	ticket := &domain.Ticket{
		Monto:         req.Monto,
		FuncionId:     req.FuncionId,
		TransaccionId: req.TransaccionId,
		EmpleadoId:    e.(*token.Payload).ID,
		Boletos:       boletos,
	}
	err := h.s.CreateTicket(ctx, ticket)
	if err != nil {
		log.Printf("Failed to create ticket: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ticketId": ticket.Id})
}

func (h *handler) IniciarCompra(c *gin.Context) {
	var req inicarCompraReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	err := h.s.IniciarCompra(ctx, req.TransaccionId, req.Boletos)
	if err != nil {
		log.Printf("Failed to get funciones: %v\n", err)
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.Status(http.StatusNoContent)
}
