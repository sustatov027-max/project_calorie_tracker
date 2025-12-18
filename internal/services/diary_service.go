package services

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/internal/repositories"
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
	if err != nil{
		return models.MealLog{}, err
	}

	return createdMeal, nil
}
