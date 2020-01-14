package jwtMiddleware

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTCustomClaims model for custom claims
type JWTCustomClaims struct {
	Email     string `json:"email"`
	TokenType string `json:"tokenType"`
	Admin     bool   `json:"admin"`
	jwt.StandardClaims
}
