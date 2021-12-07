package token

import "time"

// Token permite implementar mocks para probar los handlers
type Maker interface {
	// CreateToken crea un token para un usuario y duracion especificos
	CreateToken(id, email string, duration time.Duration) (string, error)

	// Verifica si un token es valido o no
	VerifyToken(token string) (*Payload, error)
}
