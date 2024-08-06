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
		os.Getenv("ORDER_DB_HOST"),
		os.Getenv("ORDER_DB_USERNAME"),
		os.Getenv("ORDER_DB_USER_PASSWORD"),
		os.Getenv("ORDER_DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database.ConnectPaymentDB(
		os.Getenv("PAYMENT_DB_HOST"),
		os.Getenv("PAYMENT_DB_USERNAME"),
		os.Getenv("PAYMENT_DB_USER_PASSWORD"),
		os.Getenv("PAYMENT_DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	server01 := &http.Server{
		Addr:         ":8081",
		Handler:      orderRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8082",
		Handler:      paymentRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
