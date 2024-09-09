package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtSecret = []byte("mySecretKey")

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ServiceJwtImpl struct{}

func NewJWTServiceImpl() ServiceJwt {
	return &ServiceJwtImpl{}
}

func (serviceJwt *ServiceJwtImpl) GenerateToken(username string) (string, error) {
	claims := &JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token berlaku 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (serviceJwt *ServiceJwtImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(*JwtCustomClaims); !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
