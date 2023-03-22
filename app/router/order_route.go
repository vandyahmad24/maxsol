package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"vandyahmad24/maxsol/app/domain/repository/cake_repository"
	"vandyahmad24/maxsol/app/domain/repository/order_repository"
	"vandyahmad24/maxsol/app/handler"
	"vandyahmad24/maxsol/app/usecase/order_usecase"
)

func OrderRouter(router *fiber.App, db *gorm.DB) {
	orderRepository := order_repository.NewOrder(db)
	cakeRepository := cake_repository.NewCake(db)
	orderService := order_usecase.NewOrderUsecase(orderRepository, cakeRepository)
	roleHandler := handler.NewOrderServiceService(orderService)
	router.Post("/orders", roleHandler.CreateOrderService)
	router.Get("/orders", roleHandler.GetAllOrderService)
	router.Get("/orders/:id", roleHandler.GetOrderService)
	router.Delete("/orders/:id", roleHandler.DeleteOrderService)
	router.Patch("/orders/:id", roleHandler.UpdateOrderService)
	router.Post("/orders-bulk", roleHandler.CreateBulkOrderService)
}
