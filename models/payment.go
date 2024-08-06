package models

import (
	"gorm.io/gorm"
)

type PaymentStatus string

const (
	StatusConfirmed PaymentStatus = "confirmed"
	StatusDeclined PaymentStatus = "declined"
)

type Payment struct {
	gorm.Model
	OrderID     uint   `json:"order_id"`
	Amount      float64 `json:"amount"`
	Method      string `json:"method"`
	Status      string `json:"status"`
}
