package gateways

import "github.com/dgrijalva/jwt-go"

type JwtMiddleware interface {
	CreateToken(userID string) string
	ParseToken(jwtToken string) (jwt.MapClaims, error)
	// CheckJwtToken(jwtToken string) bool
}
