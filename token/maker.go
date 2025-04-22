package token

import "time"

// Maker es una interfaz para el manejamiento de tokens
type Maker interface {
	// CreateToken crea un nuevo token con un usuario y duración en específico
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken cheka si el token es valido o no
	VerifyToken(token string) (*Payload, error)
}
