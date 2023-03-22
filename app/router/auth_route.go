package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/repository/user_repository"
	"vandyahmad24/maxsol/app/handler"
	"vandyahmad24/maxsol/app/usecase/jwt_usecase"
	"vandyahmad24/maxsol/app/usecase/user_usecase"
)

func AuthRouter(router *fiber.App, db *gorm.DB) {
	userRepository := user_repository.NewUserRepository(db)
	userService := user_usecase.NewUserUsecase(userRepository)
	authInteractor := jwt_usecase.NewServiceJwt()
	userHandler := handler.NewUserServiceService(userService, authInteractor)
	router.Post("/register", userHandler.RegisterService)
	router.Post("/login", userHandler.LoginService)
}
