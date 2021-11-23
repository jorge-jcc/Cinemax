package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type createEmpleadoReq struct {
	Nombre    string `json:"nombre" binding:"required"`
	ApPaterno string `json:"apPaterno" binding:"required"`
	ApMaterno string `json:"apMaterno"`
	Rfc       string `json:"rfc" binding:"required"`
	Edad      int8   `json:"edad" binding:"required,numeric"`
	Email     string `json:"email" binding:"required,email"`
	Direccion string `json:"direccion" binding:"required"`
	Telefono  string `json:"telefono" binding:"required,len=10"`
	Password  string `json:"password" binding:"required,gte=6,lte=30"`
}

type loginEmpleadoReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type loginRes struct {
	AccessToken string           `json:"accessToken"`
	Empleado    *domain.Empleado `json:"user"`
}

func (h *handler) CreateEmpleado(c *gin.Context) {
	// Enlaza el JSON entrante a estructura req y verificar errores de validación
	var req createEmpleadoReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	e := domain.NewEmpleado(
		req.Nombre,
		req.ApPaterno,
		req.ApMaterno,
		req.Rfc,
		req.Email,
		req.Direccion,
		req.Telefono,
		req.Password,
		req.Edad,
	)

	ctx := c.Request.Context()

	err := h.s.CreateEmpleado(ctx, e)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	t, err := h.t.CreateToken(e.Email, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginRes{AccessToken: t})
}

func (h *handler) LoginEmpleado(c *gin.Context) {
	// Enlaza el JSON entrante a estructura y verificar errores de validación
	var req loginEmpleadoReq
	if ok := util.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	e, err := h.s.LoginEmpleado(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(domain.Status(err), gin.H{
			"error": err,
		})
		return
	}
	t, err := h.t.CreateToken(e.Email, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loginRes{
		AccessToken: t,
		Empleado:    e,
	})
}
