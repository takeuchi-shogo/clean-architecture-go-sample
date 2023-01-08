package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
)

type JwtAuth struct {
	// Claims    jwt.Claims
	TokenExpireAt int
	SecretKey     string
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func NewJwtAuth(c *config.Config) *JwtAuth {
	return &JwtAuth{
		// Claims:    claim,
		TokenExpireAt: c.Jwt.TokenExpireAt,
		SecretKey:     c.Jwt.SecretKey,
	}
}

// Create json web token
func (j *JwtAuth) CreateToken(userID int) string {
	claim := jwt.MapClaims{
		"userId":   userID,
		"expireAt": time.Now().Add(time.Hour * time.Duration(j.TokenExpireAt)).Unix(), // 1day
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	// Add Signature to Token
	tokenString, _ := token.SignedString([]byte(j.SecretKey))

	return tokenString
}

func (j *JwtAuth) ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("SecretKey"), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

// Check json web token in request header
func (j *JwtAuth) CheckJwtToken(c *gin.Context) {
	token, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}
	t, err := j.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}
}
