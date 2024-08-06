package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderState string

const (
	StateCreated   OrderState = "created"
	StateConfirmed OrderState = "confirmed"
	StateDelivered OrderState = "delivered"
	StateCancelled OrderState = "cancelled"
)

type Order struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`
	Amount    float64        `json:"amount"`
	State     OrderState     `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
