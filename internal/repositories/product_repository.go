package repositories

import (
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/database"
)

type ProductRepository struct{}


func (r *ProductRepository)InsertProduct(product *models.Product) error {
	db := database.DB()
	result := db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ProductRepository) ExtractProducts() ([]models.Product, error) {
	db := database.DB()
	products := []models.Product{}

	result := db.Find(&products)
	if result.Error != nil {
		return []models.Product{}, result.Error
	}

	return products, nil
}

func (r *ProductRepository) DeleteProduct(id string) error {
	db := database.DB()

	result := db.Delete(&models.Product{}, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateProduct(product *models.Product) (models.Product, error) {
	db := database.DB()

	updateProduct := models.Product{}
	result := db.Model(&models.Product{}).Where("id = ?", product.ID).Updates(product)
	db.First(&updateProduct, product.ID)
	if err := result.Error; err != nil {
		return models.Product{}, err
	}

	return updateProduct, nil
}

func (r *ProductRepository) GetProductByID(id int) (models.Product, error){
	db := database.DB()

	getProduct := models.Product{}
	result := db.First(&getProduct, id)
	if err := result.Error; err != nil{
		return models.Product{}, err
	}
	return getProduct, nil
}
