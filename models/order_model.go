package models

import "time"

type Order struct {
	OrderId      uint        `gorm:"primary_key" json:"order_id"`
	CustomerName string      `json:"customer_name"`
	OrderedAt    time.Time   `json:"ordered_at"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Items        []OrderItem `gorm:"foreignkey:OrderId" json:"items"`
}
