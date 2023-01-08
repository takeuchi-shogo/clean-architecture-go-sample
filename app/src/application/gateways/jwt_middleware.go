package gateways

import "github.com/dgrijalva/jwt-go"

type JwtMiddleware interface {
	CreateToken(userID int) string
	ParseToken(jwtToken string) (jwt.MapClaims, error)
	// CheckJwtToken(jwtToken string) bool
}
