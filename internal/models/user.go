package models

type User struct {
	ID           int `gorm:"unique"`
	Name         string
	Age          int
	Email        string `gorm:"unique"`
	PasswordHash string `json:"-"`
	Weight       float64
	Height       float64
	Gender       string
	ActiveDays   int
	CaloriesNorm float64
}
