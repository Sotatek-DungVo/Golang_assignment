package main

import (
	"log"
	"micro/controllers"
	"micro/database"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func orderRouter() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.POST("/orders", controllers.CreateOrder)
	e.PUT("/orders/:id/cancel", controllers.CancelOrder)
	e.GET("/orders/:id/status", controllers.CheckOrderStatus)
	e.POST("/orders/delivery", controllers.DeliveryOrder)

	return e
}

func paymentRouter() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.POST("/payments", controllers.ProcessPayment)

	return e
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectOrderDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_USER_PASSWORD"),
		os.Getenv("ORDER_DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database.ConnectPaymentDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_USER_PASSWORD"),
		os.Getenv("PAYMENT_DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	orderPort := os.Getenv("ORDER_PORT")
	paymentPort := os.Getenv("PAYMENT_PORT")

	orderServer := &http.Server{
		Addr:         ":" + orderPort,
		Handler:      orderRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	paymentServer := &http.Server{
		Addr:         ":" + paymentPort,
		Handler:      paymentRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return orderServer.ListenAndServe()
	})

	g.Go(func() error {
		return paymentServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
