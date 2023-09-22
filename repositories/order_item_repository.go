package repositories

import (
	"gorm.io/gorm"
	"hacktiv8-assignment2/models"
	requests "hacktiv8-assignment2/requests/orders"
)

type OrderItemInterface interface {
	CreateOrderItem(request requests.NewOrderItemRequest) (models.OrderItem, error)
	DeleteOrderItem(orderItemId int) error
	DeleteOrderItemsByOrderId(orderId int) error
}

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return OrderItemRepository{db: db}
}

func (orderItemRepository *OrderItemRepository) CreateOrderItem(request requests.NewOrderItemRequest, orderId int) (models.OrderItem, error) {
	orderItem := models.OrderItem{
		OrderId:     uint(orderId),
		ItemCode:    request.ItemCode,
		Description: request.Description,
		Quantity:    request.Quantity,
	}
	orderItemRepository.db.Create(&orderItem)
	return orderItem, nil
}

func (orderItemRepository *OrderItemRepository) DeleteOrderItem(orderItemId int) error {
	orderItem := models.OrderItem{}
	orderItemRepository.db.Where("id = ?", orderItemId).Delete(&orderItem)
	return nil
}

func (orderItemRepository *OrderItemRepository) DeleteOrderItemsByOrderId(orderId int) error {
	orderItem := models.OrderItem{}
	orderItemRepository.db.Where("order_id = ?", orderId).Delete(&orderItem)
	return nil
}
