package services

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"time"
)

type ProductRepo interface{
	InsertProduct(product *models.Product) error
	ExtractProducts() ([]models.Product, error)
	DeleteProduct(id string) error
	UpdateProduct(product *models.Product) (models.Product, error)
	GetProductByID(id int) (models.Product, error)
}

type ProductService struct{
	postgres ProductRepo
}

func NewProductService(r ProductRepo) *ProductService{
	return &ProductService{postgres: r}
}

func (s *ProductService) CreateProduct(name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error) {
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

	err := s.postgres.InsertProduct(&newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.postgres.ExtractProducts()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.postgres.DeleteProduct(id)
}

func (s *ProductService) UpdateProduct(id int, name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error) {
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

	product := models.Product{ID: id, Name: name, Calories: calories, Proteins: proteins, Fats: fats, Carbohydrates: carbohydrates}
	return s.postgres.UpdateProduct(&product)
}

func (s *ProductService) CalculateCPFC(product models.Product, gramms float64) (float64, float64, float64, float64) {
	calories := (product.Calories / 100) * gramms
	proteins := (product.Proteins / 100) * gramms
	fats := (product.Fats / 100) * gramms
	carbohydrates := (product.Carbohydrates / 100) * gramms

	return calories, proteins, fats, carbohydrates
}

func (s *ProductService) GetProductByID(id int) (models.Product, error) {
	getProduct, err := s.postgres.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}

	return getProduct, nil
}
