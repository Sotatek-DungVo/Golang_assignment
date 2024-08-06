package dtos

type OrderResponse struct {
	Message string `json:"message"`
}

type PaymentPayload struct {
	OrderID uint `json:"orderId"`
	Status  string    `json:"status"`
}
