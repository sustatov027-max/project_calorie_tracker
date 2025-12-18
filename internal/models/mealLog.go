package models

import "time"

type MealLog struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int  `gorm:"index"`
	ProductID int
	Product   Product
	Gramms    float64
	CreatedAt time.Time
}
