package middlewares

type JwtMiddleware interface {
	CreateToken() string
	CheckJwtToken(jwtToekn string) bool
}

// func NewJwtMiddleware() *JwtMiddleware {
// 	return
// }
