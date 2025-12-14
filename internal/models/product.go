package models

import (
	"time"
)

type Product struct {
	ID int `gorm:"unique"`
	Name string
	Calories float64
	Proteins float64
	Fats float64
	Carbohydrates float64
	CreatedAt time.Time
}