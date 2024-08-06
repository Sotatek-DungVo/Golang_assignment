package controllers

import (
	"micro/dtos"
	"micro/models"
	"micro/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdOrder, err := services.CreateOrder(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdOrder)
}

func CancelOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := services.CancelOrder(id)

	if err != nil {
		if err.Error() == "Order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}

func CheckOrderStatus(c *gin.Context) {
	id := c.Param("id")
	order, err := services.GetOrderStatus(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeliveryOrder(c *gin.Context) {
	var paymentPayload dtos.PaymentPayload
	c.BindJSON(&paymentPayload)

	response := services.DeliveryOrder(paymentPayload)

	c.JSON(http.StatusOK, response)
}