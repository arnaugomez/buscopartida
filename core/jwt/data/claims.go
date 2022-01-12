package jwtData

import "github.com/dgrijalva/jwt-go"

type claims struct {
	UserId uint
	jwt.StandardClaims
}
