package services

import (
	"hacktiv8-assignment2/app/models"
	"hacktiv8-assignment2/app/repositories"
	requests "hacktiv8-assignment2/app/requests/orders"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return OrderService{
		orderRepository: orderRepository,
	}
}

func (service *OrderService) GetOrders() ([]models.Order, error) {
	return service.orderRepository.GetOrders()
}

func (service *OrderService) CreateOrder(request requests.NewOrderRequest) (models.Order, error) {
	return service.orderRepository.CreateOrder(request)
}

func (service *OrderService) GetOrderById(id int) (models.Order, error) {
	return service.orderRepository.GetOrderById(id)
}
