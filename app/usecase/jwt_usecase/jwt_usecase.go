package jwt_usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"vandyahmad24/maxsol/app/config"
	"vandyahmad24/maxsol/app/model"
)

type jwtService struct {
}

func NewServiceJwt() *jwtService {
	return &jwtService{}
}

var key string

func (s *jwtService) GenerateToken(user *model.User) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}
	key = cfg.Rest.JwtKey
	claim := jwt.MapClaims{}
	claim["user_id"] = user.Id
	claim["name"] = user.Name
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	key = cfg.Rest.JwtKey
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")

		}
		return []byte(key), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
