package jwt_usecase

import (
	"github.com/golang-jwt/jwt"
	"vandyahmad24/maxsol/app/model"
)

type AuthServicePort interface {
	GenerateToken(user *model.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
