package database

import (
	"fmt"
	"log"
	"os"

	"github.com/sustatov027-max/project_calorie_tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to database: ", err.Error())
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.MealLog{})
	return db
}

func DB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
	}
	return dbase
}
