package models

import "time"

type OrderItem struct {
	ItemId      uint `gorm:"primary_key"`
	ItemCode    string
	Description string
	Quantity    uint
	OrderId     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
