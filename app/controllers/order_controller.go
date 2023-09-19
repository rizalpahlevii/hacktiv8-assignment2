package controllers

import "github.com/gin-gonic/gin"

type OrderController struct {
}

func NewOrderController() OrderController {
	return OrderController{}
}

func (controller *OrderController) GetOrders(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Get Orders",
	})
}
