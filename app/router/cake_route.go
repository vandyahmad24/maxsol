package router

import (
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/handler"
	"vandyahmad24/maxsol/app/usecase/cake_usecase"

	"github.com/gofiber/fiber/v2"
)

func CakeRouter(router *fiber.App, db *gorm.DB) {
	cakeRepository := cake_repository.NewCake(db)
	cakeService := cake_usecase.NewCakeUsecase(cakeRepository)
	roleHandler := handler.NewCakeServiceService(cakeService)
	router.Post("/cakes", roleHandler.CreateCakeService)
	router.Get("/cakes", roleHandler.GetAllCakeService)
	router.Get("/cakes/:id", roleHandler.GetCakeService)
	router.Delete("/cakes/:id", roleHandler.DeleteCakeService)
	router.Patch("/cakes/:id", roleHandler.UpdateCakeService)
}
