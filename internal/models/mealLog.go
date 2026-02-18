package models

import "time"

type MealLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"index" json:"user_id"`
	ProductID int       `json:"product_id"`
	Product   Product   `json:"product"`
	Grams     float64   `gorm:"column:gramms" json:"grams"`
	CreatedAt time.Time `json:"created_at"`
}
