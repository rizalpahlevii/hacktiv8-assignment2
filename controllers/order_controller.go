package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/helpers"
	requests "hacktiv8-assignment2/requests/orders"
	"hacktiv8-assignment2/services"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{orderService: orderService}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object}  helpers.FormatCollectionResponse
// @Router /orders [get]
func (controller *OrderController) GetOrders(context *gin.Context) {
	orders, err := controller.orderService.GetOrders()

	if err != nil {
		helpers.ErrorResponse(context, http.StatusInternalServerError, "Failed to retrieve orders", nil)
		return
	}

	context.JSON(200, helpers.FormatCollectionResponse{
		Code: 200,
		Data: orders,
	})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body requests.NewOrderRequest true "Order"
// @Success 201 {object}  helpers.FormatResponse
// @Router /orders [post]
func (controller *OrderController) CreateOrder(context *gin.Context) {
	var newOrderRequest requests.NewOrderRequest

	if err := context.ShouldBindJSON(&newOrderRequest); err != nil {
		fmt.Println(err)
		helpers.ErrorResponse(context, http.StatusUnprocessableEntity, "invalid json request", nil)
		return
	}

	order, err := controller.orderService.CreateOrder(newOrderRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusInternalServerError, "Failed to create order", nil)
		return
	}

	helpers.SuccessResponse(context, http.StatusCreated, "Order created successfully", order)
}

// GetOrder godoc
// @Summary Get order by id
// @Description Get order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object}  helpers.FormatResponse
// @Router /orders/{id} [get]
func (controller *OrderController) GetOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	order, err := controller.orderService.GetOrderById(id)

	if err != nil {
		helpers.ErrorResponse(context, http.StatusNotFound, "Order not found", nil)
		return
	}

	helpers.SuccessResponse(context, http.StatusOK, "Order retrieved successfully", order)
}

// DeleteOrder godoc
// @Summary Delete order by id
// @Description Delete order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object}  helpers.FormatResponse
// @Router /orders/{id} [delete]
func (controller *OrderController) DeleteOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	err := controller.orderService.DeleteOrder(id)

	if err != nil {
		helpers.ErrorResponse(context, http.StatusInternalServerError, "Failed to delete order", nil)
		return
	}

	helpers.SuccessResponse(context, http.StatusOK, "Order deleted successfully", nil)
}

// UpdateOrder godoc
// @Summary Update order by id
// @Description Update order by id
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body requests.NewOrderRequest true "Order"
// @Success 200 {object}  helpers.FormatResponse
// @Router /orders/{id} [put]
func (controller *OrderController) UpdateOrder(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var newOrderRequest requests.NewOrderRequest
	if err := context.ShouldBindJSON(&newOrderRequest); err != nil {
		helpers.ErrorResponse(context, http.StatusUnprocessableEntity, "invalid json request", nil)
		return
	}

	order, err := controller.orderService.UpdateOrder(id, newOrderRequest)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	helpers.SuccessResponse(context, http.StatusOK, "Order updated successfully", order)
}
