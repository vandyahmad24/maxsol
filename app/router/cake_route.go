package router

import (
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/handler"
	"vandyahmad24/maxsol/app/middleware"
	"vandyahmad24/maxsol/app/usecase/cake_usecase"

	"github.com/gofiber/fiber/v2"
)

func CakeRouter(router *fiber.App, db *gorm.DB) {
	cakeRepository := cake_repository.NewCake(db)
	cakeService := cake_usecase.NewCakeUsecase(cakeRepository)
	roleHandler := handler.NewCakeServiceService(cakeService)
	api := router.Group("", middleware.JWTProtected)
	api.Post("/cakes", roleHandler.CreateCakeService)
	api.Get("/cakes", roleHandler.GetAllCakeService)
	api.Get("/cakes/:id", roleHandler.GetCakeService)
	api.Delete("/cakes/:id", roleHandler.DeleteCakeService)
	api.Patch("/cakes/:id", roleHandler.UpdateCakeService)
}
