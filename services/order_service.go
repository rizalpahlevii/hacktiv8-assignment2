package services

import (
	"gorm.io/gorm"
	"hacktiv8-assignment2/models"
	"hacktiv8-assignment2/repositories"
	requests "hacktiv8-assignment2/requests/orders"
)

type OrderService struct {
	db                  *gorm.DB
	orderRepository     repositories.OrderRepository
	orderItemRepository repositories.OrderItemRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, orderItemRepository repositories.OrderItemRepository) OrderService {
	return OrderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
	}
}

func (service *OrderService) GetOrders() ([]models.Order, error) {
	return service.orderRepository.GetOrders()
}

func (service *OrderService) CreateOrder(request requests.NewOrderRequest) (models.Order, error) {
	order, err := service.orderRepository.CreateOrder(request)
	if err != nil {
		return order, err
	}

	for _, orderItemRequest := range request.Items {
		item := requests.NewOrderItemRequest{
			ItemCode:    orderItemRequest.ItemCode,
			Description: orderItemRequest.Description,
			Quantity:    orderItemRequest.Quantity,
		}
		_, err := service.orderItemRepository.CreateOrderItem(item, int(order.OrderId))
		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}

func (service *OrderService) GetOrderById(id int) (models.Order, error) {
	return service.orderRepository.GetOrderById(id)
}

func (service *OrderService) DeleteOrder(id int) error {
	order, err := service.orderRepository.GetOrderById(id)
	if err != nil {
		return err
	}
	err = service.orderItemRepository.DeleteOrderItemsByOrderId(int(order.OrderId))
	if err != nil {
		return err
	}
	return service.orderRepository.DeleteOrder(order)
}

func (service *OrderService) UpdateOrder(id int, request requests.NewOrderRequest) (models.Order, error) {

	order, err := service.orderRepository.GetOrderById(id)
	if err != nil {
		return models.Order{}, err
	}
	order, err = service.orderRepository.UpdateOrder(order, request)
	if err != nil {
		return models.Order{}, err
	}

	err = service.orderItemRepository.DeleteOrderItemsByOrderId(int(order.OrderId))
	if err != nil {
		return models.Order{}, err
	}

	for _, orderItemRequest := range request.Items {
		item := requests.NewOrderItemRequest{
			ItemCode:    orderItemRequest.ItemCode,
			Description: orderItemRequest.Description,
			Quantity:    orderItemRequest.Quantity,
		}
		_, err := service.orderItemRepository.CreateOrderItem(item, int(order.OrderId))
		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}
