package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	costInt, err := strconv.Atoi(os.Getenv("COST"))
	if err != nil{
		return "", err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), costInt)
	if err != nil{
		return "", err
	}

	return string(passwordHash), nil
}