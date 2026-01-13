package utils

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	costStr := os.Getenv("COST")
	costStr = strings.TrimSpace(costStr)

	costInt, err := strconv.Atoi(costStr)
	if err != nil {
		return "", err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), costInt)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
