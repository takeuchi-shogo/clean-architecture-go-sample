package gateways

type JwtMiddleware interface {
	CheckJwtToken(jwtToken string) bool
}
