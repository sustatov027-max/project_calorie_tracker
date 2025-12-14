package services

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/internal/repositories"
	"time"
)

func CreateProduct(name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error) {
	if name == "" {
		return models.Product{}, errors.New("product name is required")
	}
	if calories < 0 {
		return models.Product{}, errors.New("calories must be >= 0")
	}
	if proteins < 0 {
		return models.Product{}, errors.New("proteins must be >= 0")
	}
	if fats < 0 {
		return models.Product{}, errors.New("fats must be >= 0")
	}
	if carbohydrates < 0 {
		return models.Product{}, errors.New("carbohydrates must be >= 0")
	}

	newProduct := models.Product{ID: 0, Name: name, Calories: calories, Proteins: proteins, Fats: fats, Carbohydrates: carbohydrates, CreatedAt: time.Now().Local()}

	err := repositories.InsertProduct(&newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

func GetAllProducts() ([]models.Product, error) {
	products, err := repositories.ExtractProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func DeleteProduct(id string) error {
	return repositories.DeleteProduct(id)
}

func UpdateProduct(id int, name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error) {
	product := models.Product{ID: id, Name: name, Calories: calories, Proteins: proteins, Fats: fats, Carbohydrates: carbohydrates}
	return repositories.UpdateProduct(&product)
}
