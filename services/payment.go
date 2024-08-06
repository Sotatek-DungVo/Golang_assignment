package services

import (
	"math/rand"
	"micro/models"
	"micro/repositories"
)

func ProcessPayment(payment models.Payment) (models.Payment, error) {
	if rand.Float32() < 0.5 {
		payment.Status = string(models.StatusDeclined)
	} else {
		payment.Status = string(models.StatusConfirmed)
	}

	if err := repositories.CreatePayment(&payment); err != nil {
		return payment, err
	}

	return payment, nil
}
