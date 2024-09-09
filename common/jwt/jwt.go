package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type ServiceJwt interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
