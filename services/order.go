package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"micro/database"
	"micro/dtos"
	"micro/models"
	"micro/repositories"
	"net/http"
	"os"
	"time"
)

type PaymentResponse struct {
	CreatedAt string  `json:"CreatedAt"`
	DeletedAt *string `json:"DeletedAt"`
	ID        int     `json:"ID"`
	UpdatedAt string  `json:"UpdatedAt"`
	Amount    int     `json:"amount"`
	Method    string  `json:"method"`
	OrderID   int     `json:"order_id"`
	Status    string  `json:"status"`
}

func CreateOrder(order models.Order) (models.Order, error) {
	order.State = models.StateCreated

	if err := repositories.CreateOrder(&order); err != nil {
		return order, err
	}

	paymentPayload := map[string]interface{}{
		"order_id": order.ID,
		"amount":   order.Amount,
		"method":   "dummy_method",
	}

	paymentPayloadJSON, err := json.Marshal(paymentPayload)

	if err != nil {
		return order, errors.New("failed to create payment payload")
	}

	paymentURL := os.Getenv("PAYMENTS_SERVICE_URL") + "/payments"
	response, err := http.Post(paymentURL, "application/json", bytes.NewBuffer(paymentPayloadJSON))

	if err != nil || response.StatusCode != http.StatusOK {
		order.State = models.StateCancelled
		repositories.UpdateOrder(&order)
		return order, errors.New("payment processing failed")
	}

	defer response.Body.Close()

	var paymentRes PaymentResponse
	json.NewDecoder(response.Body).Decode(&paymentRes)

	if paymentRes.Status == string(models.StatusConfirmed) {
		order.State = models.StateConfirmed
	} else {
		order.State = models.StateCancelled
	}

	if err := repositories.UpdateOrder(&order); err != nil {
		return order, err
	}

	return order, nil
}

func CancelOrder(id string) (models.Order, error) {
	order, err := repositories.GetOrderById(id)

	if err != nil {
		return order, errors.New("order not found")
	}

	if order.State == models.StateCancelled {
		return order, errors.New("this order is already cancelled")
	}

	if order.State == models.StateDelivered {
		return order, errors.New("this order is already delivered")
	}

	order.State = models.StateCancelled

	if err := repositories.UpdateOrder(&order); err != nil {
		return order, err
	}

	return order, nil
}

func GetOrderStatus(id string) (models.Order, error) {
	return repositories.GetOrderById(id)
}

func DeliveryOrder(paymentPayload dtos.PaymentPayload) dtos.OrderResponse {
	if paymentPayload.Status == string(models.StatusConfirmed) {
		time.AfterFunc(10*time.Second, func() {
			var order models.Order
			database.OrderDB.Where("id = ?", paymentPayload.OrderID).First(&order)

			if order.State == models.StateConfirmed {
				order.State = models.StateDelivered
				database.OrderDB.Save(&order)
			}
		})
	}

	return dtos.OrderResponse{Message: "Order updated"}
}
