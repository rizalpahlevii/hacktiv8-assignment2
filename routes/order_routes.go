package routes

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/config"
	"hacktiv8-assignment2/controllers"
	"hacktiv8-assignment2/repositories"
	"hacktiv8-assignment2/services"
)

func OrderRoutes(router *gin.Engine) {
	db := config.GetDatabaseConnection()
	orderRepository := repositories.NewOrderRepository(db)
	orderItemRepository := repositories.NewOrderItemRepository(db)
	orderService := services.NewOrderService(orderRepository, orderItemRepository)
	orderController := controllers.NewOrderController(orderService)

	router.GET("/orders", orderController.GetOrders)
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:id", orderController.GetOrder)
	router.PUT("/orders/:id", orderController.UpdateOrder)
	router.DELETE("/orders/:id", orderController.DeleteOrder)
}
