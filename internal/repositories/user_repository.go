package repositories

import (
	"errors"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/database"
)

func SaveUser(user *models.User) error{
	db := database.Init()

	result := db.Create(user)
	if err := result.Error; err != nil{
		return err
	}

	return nil
}

func ExtractUser(email string) (models.User, error){
	db := database.Init()

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

func GetUserByID(userID any) (models.User, error){
	db := database.Init()

	var user models.User
	db.First(&user, userID)
	if user.ID == 0{
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}