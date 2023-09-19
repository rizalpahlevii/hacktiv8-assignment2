package controllers

import (
	"github.com/gin-gonic/gin"
	requests "hacktiv8-assignment2/app/requests/orders"
	"hacktiv8-assignment2/app/services"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{orderService: orderService}
}

func (controller *OrderController) GetOrders(context *gin.Context) {
	orders, err := controller.orderService.GetOrders()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"orders": orders})
}

func (controller *OrderController) CreateOrder(context *gin.Context) {
	var newOrderRequest requests.NewOrderRequest

	if err := context.ShouldBindJSON(&newOrderRequest); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid json request"})
		return
	}

	order, err := controller.orderService.CreateOrder(newOrderRequest)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

func (controller *OrderController) GetOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	order, err := controller.orderService.GetOrderById(id)

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"order": order})
}
