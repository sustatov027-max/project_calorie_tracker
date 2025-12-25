package services

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/internal/repositories"
	"strconv"
	"time"
)

type DiaryRepo interface{
	InsertMeal(mealLog *models.MealLog) error
	ExtractMeals(userID int, date time.Time) ([]models.MealLog, error)
	DeleteMeal(userID int, id string) error
	UpdateMeal(userID int, meal *models.MealLog) (models.MealLog, error)
}

type DiaryService struct{
	postgres DiaryRepo
}

func NewDiaryService(r DiaryRepo) *DiaryService{
	return &DiaryService{postgres: r}
}

var productService = NewProductService(&repositories.ProductRepository{})
var userService = NewUserService(&repositories.UserRepository{})

func (s *DiaryService) CreateMeal(userID int, productID int, gramms float64) (models.MealLog, error) {
	if gramms <= 0 {
		return models.MealLog{}, errors.New("gramms must be greater than 0")
	}

	productByID, err := productService.GetProductByID(productID)
	if err != nil {
		return models.MealLog{}, err
	}
	createdMeal := models.MealLog{UserID: userID, ProductID: productID, Product: productByID, Gramms: gramms, CreatedAt: time.Now()}

	err = s.postgres.InsertMeal(&createdMeal)
	if err != nil {
		return models.MealLog{}, err
	}

	return createdMeal, nil
}

func (s *DiaryService) GetAllMealsForDay(userID int, date time.Time) ([]models.MealLog, error) {
	meals, err := s.postgres.ExtractMeals(userID, date)
	if err != nil {
		return []models.MealLog{}, err
	}

	return meals, nil
}

func (s *DiaryService) DeleteMeal(userID int, id string) error {
	return s.postgres.DeleteMeal(userID, id)
}

func (s *DiaryService) UpdateMeal(userID int, id string, productID int, gramms float64) (models.MealLog, error) {
	product, err := productService.GetProductByID(productID)
	if err != nil {
		return models.MealLog{}, err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.MealLog{}, err
	}
	updatedMeal := models.MealLog{ID: uint(idInt), ProductID: productID, Product: product, Gramms: gramms}

	return s.postgres.UpdateMeal(userID, &updatedMeal)
}

func (s *DiaryService) Summary(userID int) (models.DaySummary, error) {
	meals, err := s.GetAllMealsForDay(userID, time.Now().Local())
	if err != nil {
		return models.DaySummary{}, err
	}

	var daySummary models.DaySummary

	for _, meal := range meals {
		calories, proteins, fats, carbohydrates := productService.CalculateCPFC(meal.Product, meal.Gramms)

		daySummary.TotalCalories += calories
		daySummary.TotalProteins += proteins
		daySummary.TotalFats += fats
		daySummary.TotalCarbs += carbohydrates
	}

	daySummary.Meals = meals

	user, err := userService.GetUser(userID)
	if err != nil {
		return models.DaySummary{}, err
	}

	daySummary.DailyNorm = user.CaloriesNorm
	daySummary.Remaining = user.CaloriesNorm - daySummary.TotalCalories

	return daySummary, nil
}
