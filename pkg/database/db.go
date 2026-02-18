package database

import (
	"fmt"
	"log"

	"github.com/sustatov027-max/project_calorie_tracker/internal/config"
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	if dbase != nil {
		return dbase
	}

	cfg := config.MustGet()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to database: ", err.Error())
	}

	if err := db.AutoMigrate(&models.Product{}, &models.User{}, &models.MealLog{}); err != nil {
		log.Fatal("Error while migration: ", err.Error())
	}

	dbase = db
	return dbase
}

func DB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
	}
	return dbase
}
