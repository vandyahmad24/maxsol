package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"vandyahmad24/maxsol/app/config"
	"vandyahmad24/maxsol/app/domain/entity"
	"vandyahmad24/maxsol/app/util"
)

func JWTProtected(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(util.ApiErrorResponse("Authorization header is required", fiber.StatusUnauthorized))
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	cfg, _ := config.LoadConfig()

	key := cfg.Rest.JwtKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")

		}
		return []byte(key), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(util.ApiResponse("Invalid Token", fiber.StatusUnauthorized))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Expires time.
		userId := claims["user_id"].(float64)
		name := claims["name"].(string)
		expires := int64(claims["exp"].(float64))

		metaData := &entity.TokenMetadata{
			UserId:  userId,
			Expires: expires,
			Name:    name,
		}

		c.Locals("metaData", metaData)
	}

	return c.Next()
}
