package database

import (
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
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
