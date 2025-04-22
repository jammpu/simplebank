package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("invalid token")
)

// Payload contiene datos de el token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expire_at"`
}

// NewPayload crea un nuevo token payload con un usuario y duración especifico
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checa si el payload del token es valido o no
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// --- Implementación de jwt.Claims ---

// GetExpirationTime implementa jwt.Claims. Retorna el tiempo de expiración.
func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.ExpiredAt), nil
}

// GetIssuedAt implementa jwt.Claims. Retorna el tiempo de emisión.
func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.IssuedAt), nil
}

// GetNotBefore implementa jwt.Claims. Retorna el tiempo "no antes de".
func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(p.IssuedAt), nil
}

// GetIssuer implementa jwt.Claims. Retorna el emisor.
func (p *Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetSubject implementa jwt.Claims. Retorna el sujeto (a quién se refiere el token).
func (p *Payload) GetSubject() (string, error) {
	return "", nil
}

// GetAudience implementa jwt.Claims. Retorna la audiencia (quién debe aceptar el token).
func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
