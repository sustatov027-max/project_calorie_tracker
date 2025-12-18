package repositories

import (
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/database"
)

func InsertMeal(mealLog *models.MealLog) error{
	db := database.DB()

	result := db.Create(mealLog)
	if err := result.Error; err != nil{
		return err
	}

	return nil
}