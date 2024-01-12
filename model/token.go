package model

import jwt "github.com/golang-jwt/jwt/v4"

type Token struct {
	Token string `json:"token" binding:"required"`
}

type AccessTokenData struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenData struct {
	ID string `json:"id"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
