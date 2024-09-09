package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"task/common/jwt"
)

// JWTMiddleware membuat middleware untuk memvalidasi token JWT
func JWTMiddleware(jwtService jwt.ServiceJwt) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Ambil token dari header Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "Missing token")
			}

			// Token harus diawali dengan "Bearer "
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, "Invalid token format")
			}

			// Validasi token
			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
			}

			// Jika token valid, simpan klaim ke dalam context
			claims, ok := token.Claims.(*jwt.JwtCustomClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, "Invalid token claims")
			}

			// Menyimpan informasi pengguna ke context
			c.Set("username", claims.Username)

			// Melanjutkan ke handler berikutnya
			return next(c)
		}
	}
}
