package routes

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/app/controllers"
	"hacktiv8-assignment2/app/repositories"
	"hacktiv8-assignment2/app/services"
	"hacktiv8-assignment2/config"
)

func OrderRoutes(router *gin.Engine) {
	db := config.GetDatabaseConnection()
	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	router.GET("/orders", orderController.GetOrders)
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:id", orderController.GetOrder)
}
