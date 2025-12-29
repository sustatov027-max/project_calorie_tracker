package services

import (
	"errors"
	"math"
	"os"
	"project_calorie_tracker/internal/models"
	"project_calorie_tracker/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	SaveUser(user *models.User) error
	ExtractUser(email string) (models.User, error)
	GetUserByID(userID any) (models.User, error)
}
type UserService struct {
	postgres UserRepo
}

func NewUserService(r UserRepo) *UserService {
	return &UserService{postgres: r}
}

func (s *UserService) RegisterUser(name string, age int, email string, password string, weight float64, height float64, gender string, activeDays int) (models.User, error) {
	if name == "" {
		return models.User{}, errors.New("name must not be empty")
	}
	if age <= 0 || age > 110 {
		return models.User{}, errors.New("age must be > 0 and <= 110")
	}
	if email == "" {
		return models.User{}, errors.New("email must not be empty")
	}
	if password == "" {
		return models.User{}, errors.New("password must not be empty")
	}
	if weight <= 0 {
		return models.User{}, errors.New("weight must be greater than 0")
	}
	if height <= 0 {
		return models.User{}, errors.New("height must be greater than 0")
	}

	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	var coefficient float64

	switch activeDays {
	case 0:
		coefficient = 1.2
	case 1:
		coefficient = 1.375
	case 2:
		coefficient = 1.375
	case 3:
		coefficient = 1.55
	case 4:
		coefficient = 1.55
	case 5:
		coefficient = 1.7
	case 6:
		coefficient = 1.8
	case 7:
		coefficient = 1.9
	default:
		return models.User{}, errors.New("activeDays must be >= 0 and <= 7")
	}

	var caloriesNorm float64

	switch gender {
	case "male":
		caloriesNorm = math.Round((66.5+(13.75*weight)+(5.003*height)-(6.775*float64(age)))*coefficient*100) / 100
	case "female":
		caloriesNorm = math.Round((655.1+(9.563*weight)+(1.85*height)-(4.676*float64(age)))*coefficient*100) / 100
	default:
		return models.User{}, errors.New("gender must be male or female")
	}

	newUser := models.User{Name: name, Age: age, Email: email, PasswordHash: passwordHash, Weight: weight, Height: height, Gender: gender, ActiveDays: activeDays, CaloriesNorm: caloriesNorm}

	err = s.postgres.SaveUser(&newUser)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (s *UserService) LoginUser(email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("invalid email or password")
	}
	user, err := s.postgres.ExtractUser(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("error to create token")
	}

	return tokenString, nil
}

func (s *UserService) GetUser(userID any) (models.User, error) {
	user, err := s.postgres.GetUserByID(userID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
