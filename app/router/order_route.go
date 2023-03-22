package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/domain/repository/order_repository"
	"vandyahmad24/maxsol/app/handler"
	"vandyahmad24/maxsol/app/middleware"
	"vandyahmad24/maxsol/app/usecase/order_usecase"
)

func OrderRouter(router *fiber.App, db *gorm.DB) {
	orderRepository := order_repository.NewOrder(db)
	cakeRepository := cake_repository.NewCake(db)
	orderService := order_usecase.NewOrderUsecase(orderRepository, cakeRepository)
	roleHandler := handler.NewOrderServiceService(orderService)
	api := router.Group("", middleware.JWTProtected)
	api.Post("/orders", roleHandler.CreateOrderService)
	api.Get("/orders", roleHandler.GetAllOrderService)
	api.Get("/orders/:id", roleHandler.GetOrderService)
	api.Delete("/orders/:id", roleHandler.DeleteOrderService)
	api.Patch("/orders/:id", roleHandler.UpdateOrderService)
	api.Post("/orders-bulk", roleHandler.CreateBulkOrderService)
}
