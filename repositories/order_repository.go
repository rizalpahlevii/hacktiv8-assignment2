package repositories

import (
	"gorm.io/gorm"
	"hacktiv8-assignment2/models"
	requests "hacktiv8-assignment2/requests/orders"
)

type OrderRepositoryInterface interface {
	GetOrders() ([]models.Order, error)
	CreateOrder(request requests.NewOrderRequest) (models.Order, error)
	GetOrderById(id int) (models.Order, error)
	UpdateOrder(order models.Order, request requests.NewOrderRequest) (models.Order, error)
	DeleteOrder(order models.Order) error
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{db: db}
}

func (repository *OrderRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := repository.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (repository *OrderRepository) CreateOrder(request requests.NewOrderRequest) (models.Order, error) {
	order := models.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    request.OrderedAt,
	}
	repository.db.Create(&order)
	return order, nil
}

func (repository *OrderRepository) GetOrderById(id int) (models.Order, error) {
	var order models.Order
	result := repository.db.Preload("Items").First(&order, id)
	return order, result.Error
}

func (repository *OrderRepository) UpdateOrder(order models.Order, request requests.NewOrderRequest) (models.Order, error) {
	order.CustomerName = request.CustomerName
	order.OrderedAt = request.OrderedAt
	repository.db.Save(&order)
	return order, nil
}

func (repository *OrderRepository) DeleteOrder(order models.Order) error {
	repository.db.Delete(&order)
	return nil
}
