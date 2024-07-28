package middlewares

import "github.com/dgrijalva/jwt-go"

type Jwt interface {
	CreateToken(userID int) string
	ParseToken(jwtToken string) (jwt.MapClaims, error)
	// CheckJwtToken(jwtToken string) bool
}

type JwtMiddleware struct {
	Jwt Jwt
}

func (m *JwtMiddleware) CreateToken(userID int) string {
	return m.Jwt.CreateToken(userID)
}

func (m *JwtMiddleware) ParseToken(jwtToken string) (jwt.MapClaims, error) {
	return m.Jwt.ParseToken(jwtToken)
}
