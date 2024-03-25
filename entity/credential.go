package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

type Credential struct {
	Email    string `json:"credentialType"`
	Password string `json:"password"`
}

type JWTPayload struct {
	Id    string
	Email string
	Name  string
}

type JWTClaims struct {
	Id    string
	Email string
	Name  string
	jwt.RegisteredClaims
}
