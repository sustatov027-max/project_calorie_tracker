package database

import (
	"log"
	"os"
	"project_calorie_tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil{
		log.Fatal("Error connect to database: ", err.Error())
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
	return db
}

func DB() *gorm.DB{
	if dbase == nil{
		dbase = Init()
	}
	return dbase
}