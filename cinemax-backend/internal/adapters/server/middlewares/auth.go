package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util/token"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/domain"
)

type authHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

// ValidateJWT is a middleware that validates a jwt
func Auth(t token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := authHeader{}
		if err := c.ShouldBindHeader(&a); err != nil {
			invalidArgs := util.GetErrorsOfBinding(err)
			c.JSON(http.StatusBadRequest,
				gin.H{
					"error":       "Invalid request parameters",
					"invalidArgs": invalidArgs})
			c.Abort()
			return
		}
		idTokenHeader := strings.Split(a.Authorization, "Bearer ")
		if len(idTokenHeader) < 2 {
			err := domain.NewAuthorizationError("Must provide Authorization header with format `Bearer {token}`")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		e, err := t.VerifyToken(idTokenHeader[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("empleado", e)
		c.Next()
	}
}
