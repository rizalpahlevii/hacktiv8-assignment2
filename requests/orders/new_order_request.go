package requests

import "time"

type NewOrderRequest struct {
	CustomerName string                `json:"customer_name"`
	OrderedAt    time.Time             `json:"ordered_at"`
	Items        []NewOrderItemRequest `json:"items"`
}
