package auth

import (
	"test-fiber/dto"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	UserId int    `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func generateToken(authReq *dto.AuthRequest) (string, error) {
	claims := AuthClaims{}
}
