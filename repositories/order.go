package repositories

import (
	"micro/database"
	"micro/models"
)

func CreateOrder(order *models.Order) error {
	return database.OrderDB.Create(order).Error
}

func UpdateOrder(order *models.Order) error {
	return database.OrderDB.Save(order).Error
}

func GetOrderById(id string) (models.Order, error) {
	var order models.Order

	if err := database.OrderDB.First(&order, id).Error; err != nil {
		return order, err
	}
	
	return order, nil
}
