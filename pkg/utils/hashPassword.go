package utils

import (
	"github.com/sustatov027-max/project_calorie_tracker/internal/config"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost := config.MustGet().Cost

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
