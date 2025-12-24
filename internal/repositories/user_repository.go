package repositories

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/database"
)

type UserRepository struct{}

func (r *UserRepository) SaveUser(user *models.User) error{
	db := database.DB()

	result := db.Create(user)
	if err := result.Error; err != nil{
		return err
	}

	return nil
}

func (r *UserRepository) ExtractUser(email string) (models.User, error){
	db := database.DB()

	var user models.User
	result := db.First(&user, "email = ?", email)
	if err := result.Error; err != nil{
		return models.User{}, err
	}
	if user.ID == 0{
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(userID any) (models.User, error){
	db := database.DB()

	var user models.User
	result := db.First(&user, userID)
	if err := result.Error; err != nil{
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}