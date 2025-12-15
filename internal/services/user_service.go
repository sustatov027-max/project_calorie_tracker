package services

import (
	"errors"
	"math"
	"os"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/internal/repositories"
	"project_calorie_tracker/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name string, age int, email string, password string, weight float64, height float64, gender string, activeDays int) (models.User, error){
	passwordHash, err := utils.HashPassword(password)
	if err != nil{
		return models.User{}, err
	}

	coefficient := map[int]float64{
		0: 1.2,
		1: 1.375,
		2: 1.375,
		3: 1.55,
		4: 1.55,
		5: 1.7,
		6: 1.8,
		7: 1.9,
	}

	var caloriesNorm float64

	switch gender {
		case "male":
			caloriesNorm = math.Round((66.5 + (13.75 * weight) + (5.003 * height) - (6.775 * float64(age))) * coefficient[activeDays] * 100) / 100
		case "female":
			caloriesNorm = math.Round((655.1 + (9.563 * weight) + (1.85 * height) - (4.676 * float64(age))) * coefficient[activeDays] * 100) / 100
		default:
			return models.User{}, errors.New("gender must be male or female")
	}

	newUser := models.User{Name:name, Age:age, Email:email, PasswordHash: passwordHash, Weight: weight, Height: height, Gender: gender, ActiveDays: activeDays, CaloriesNorm: caloriesNorm}

	err = repositories.SaveUser(&newUser)
	if err != nil{
		return models.User{}, err
	}

	return newUser, nil
}

func LoginUser(email string, password string) (string, error){
	
	user, err := repositories.ExtractUser(email)
	if err != nil{
		return "", err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil{
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		return "", errors.New("error to create token")
	}

	return tokenString, nil
}

func GetUser(userID any) (models.User, error){
	user, err := repositories.GetUserByID(userID)
	if err != nil{
		return models.User{}, err
	}

	return user, nil
}
