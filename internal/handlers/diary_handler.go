package handlers

import (
	"net/http"
	"project_calorie_tracker/internal/middlewares"
	"project_calorie_tracker/internal/services"
	"project_calorie_tracker/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RegisterDiaryRoutes(r *gin.Engine) {
	r.POST("/diary", middlewares.AuthMiddleware, CreateMeal)
}

func CreateMeal(ctx *gin.Context) {
	type requestBody struct {
		ProductID int     `json:"product_id"`
		Gramms    float64 `json:"gramms"`
	}

	var body requestBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"error":"error read request body"})
		return
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error":err.Error()})
		return
	}

	createdMeal, err := services.CreateMeal(userID, body.ProductID, body.Gramms)
	if err != nil{
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error":err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdMeal)
}
