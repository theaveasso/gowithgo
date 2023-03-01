package database

import (
	"fmt"
	"log"
	"os"

	"github.com/theaveasso/goWithGO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var DB DBinstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Bangkork",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode((logger.Info)),
	})
	if err != nil {
		log.Fatal("Fail to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("runing migration")
	db.AutoMigrate(&models.Fact{})
}
