package database

import (
	"fmt"
	"log"
	"micro/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var OrderDB *gorm.DB

func ConnectOrderDB(
	DBHost string,
	DBUserName string,
	DBUserPassword string,
	DBName string,
	DBPort string,
) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DBHost, DBUserName, DBUserPassword, DBName, DBPort)

	fmt.Println(dsn)

	OrderDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	OrderDB.AutoMigrate(&models.Order{})

	fmt.Println("🚀 Connected Successfully to the Database")
}