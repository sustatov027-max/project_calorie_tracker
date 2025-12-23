package repositories

import (
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/database"
	"time"
)

func InsertMeal(mealLog *models.MealLog) error {
	db := database.DB()

	result := db.Create(mealLog)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func ExtractMeals(userID int, date time.Time) ([]models.MealLog, error) {
	db := database.DB()
	meals := []models.MealLog{}

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	result := db.Preload("Product").Where("user_id = ? AND created_at >= ? AND created_at < ?", userID, startOfDay, endOfDay).Find(&meals)
	if err := result.Error; err != nil {
		return []models.MealLog{}, err
	}

	return meals, nil
}

func DeleteMeal(userID int, id string) error {
	db := database.DB()

	result := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.MealLog{})
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func UpdateMeal(userID int, meal *models.MealLog) (models.MealLog, error) {
	db := database.DB()
	updatedMeal := models.MealLog{}

	result := db.Model(&models.MealLog{}).Where("id = ? AND user_id = ?", meal.ID, userID).Updates(meal).First(&updatedMeal, meal.ID)
	db.Preload("Product").First(&updatedMeal, meal.ID)
	if err := result.Error; err != nil {
		return models.MealLog{}, err
	}

	return updatedMeal, nil
}
