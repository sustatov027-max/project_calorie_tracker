package services

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/internal/repositories"
	"strconv"
	"time"
)

func CreateMeal(userID int, productID int, gramms float64) (models.MealLog, error) {
	if gramms <= 0 {
		return models.MealLog{}, errors.New("gramms must be greater than 0")
	}

	productByID, err := GetProductByID(productID)
	if err != nil {
		return models.MealLog{}, err
	}
	createdMeal := models.MealLog{UserID: userID, ProductID: productID, Product: productByID, Gramms: gramms, CreatedAt: time.Now()}

	err = repositories.InsertMeal(&createdMeal)
	if err != nil {
		return models.MealLog{}, err
	}

	return createdMeal, nil
}

func GetAllMealsForDay(userID int, date time.Time) ([]models.MealLog, error) {
	meals, err := repositories.ExtractMeals(userID, date)
	if err != nil {
		return []models.MealLog{}, err
	}

	return meals, nil
}

func DeleteMeal(userID int, id string) error {
	return repositories.DeleteMeal(userID, id)
}

func UpdateMeal(userID int, id string, productID int, gramms float64) (models.MealLog, error) {
	product, err := GetProductByID(productID)
	if err != nil {
		return models.MealLog{}, err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.MealLog{}, err
	}
	updatedMeal := models.MealLog{ID: uint(idInt), ProductID: productID, Product: product, Gramms: gramms}

	return repositories.UpdateMeal(userID, &updatedMeal)
}

func Summary(userID int) (models.DaySummary, error) {
	meals, err := GetAllMealsForDay(userID, time.Now().Local())
	if err != nil {
		return models.DaySummary{}, err
	}

	var daySummary models.DaySummary

	for _, meal := range meals {
		calories, proteins, fats, carbohydrates := CalculateCPFC(meal.Product, meal.Gramms)

		daySummary.TotalCalories += calories
		daySummary.TotalProteins += proteins
		daySummary.TotalFats += fats
		daySummary.TotalCarbs += carbohydrates
	}

	daySummary.Meals = meals

	service := NewUserService(&repositories.UserRepository{})
	user, err := service.GetUser(userID)
	if err != nil {
		return models.DaySummary{}, err
	}

	daySummary.DailyNorm = user.CaloriesNorm
	daySummary.Remaining = user.CaloriesNorm - daySummary.TotalCalories

	return daySummary, nil
}
