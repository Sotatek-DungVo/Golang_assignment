package controllers

import (
	"micro/models"
	"micro/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessPayment(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processedPayment, err := services.ProcessPayment(payment)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, processedPayment)
}
