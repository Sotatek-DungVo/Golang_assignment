package repositories

import (
	"micro/database"
	"micro/models"
)

func CreatePayment(payment *models.Payment) error {
	return database.PaymentDB.Create(payment).Error
}
