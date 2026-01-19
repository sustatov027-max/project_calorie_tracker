package models

type User struct {
	ID           int `gorm:"unique"`
	Name         string
	Age          int    `gorm:"type:integer;not null"`
	Email        string `gorm:"unique"`
	PasswordHash string `json:"-"`
	Weight       float64
	Height       float64
	Gender       string
	ActiveDays   int     `gorm:"type:integer"`
	CaloriesNorm float64 `gorm:"type:decimal(10,2)"`
}
