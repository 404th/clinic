package model

import jwt "github.com/golang-jwt/jwt/v4"

type Token struct {
	Token string `json:"token" binding:"required"`
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
