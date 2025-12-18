package utils

import (
    "errors"

    "github.com/gin-gonic/gin"
)

func GetUserID(ctx *gin.Context) (int, error) {
    rawID, exists := ctx.Get("userID")
    if !exists {
        return 0, errors.New("user id not found in context")
    }

    idFloat, ok := rawID.(float64)
    if !ok {
        return 0, errors.New("invalid user id type")
    }

    return int(idFloat), nil
}