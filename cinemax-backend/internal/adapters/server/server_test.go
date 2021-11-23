package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestInvalidPath(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	NewHandler(router, nil, nil)

	// Al relizar una solicitud invalida
	req := httptest.NewRequest("GET", "/invalid/123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	// El codigo de respuesta debe ser http.StatusNotFound
	require.Equal(t, resp.Code, http.StatusNotFound)
}
