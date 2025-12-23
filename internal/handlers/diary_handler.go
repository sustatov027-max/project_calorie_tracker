package handlers

import (
	"net/http"
	"project_calorie_tracker/internal/middlewares"
	"project_calorie_tracker/internal/services"
	"project_calorie_tracker/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterDiaryRoutes(r *gin.Engine) {
	r.POST("/diary", middlewares.AuthMiddleware, CreateMeal)
	r.GET("/diary", middlewares.AuthMiddleware, GetMealsForDay)
	r.DELETE("/diary/:id", middlewares.AuthMiddleware, DeleteMeal)
	r.PUT("/diary/:id", middlewares.AuthMiddleware, UpdateMeal)
	r.GET("/diary/summary", middlewares.AuthMiddleware, Summary)
}

type requestBody struct {
	ProductID int     `json:"product_id"`
	Gramms    float64 `json:"gramms"`
}

func CreateMeal(ctx *gin.Context) {
	var body requestBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"error": "error read request body"})
		return
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	createdMeal, err := services.CreateMeal(userID, body.ProductID, body.Gramms)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdMeal)
}

func GetMealsForDay(ctx *gin.Context) {
	dateStr := ctx.Query("date")

	var date time.Time
	var err error

	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"error": "invalid date"})
			return
		}
	}

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	meals, err := services.GetAllMealsForDay(userID, date)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, meals)
}

func DeleteMeal(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	err = services.DeleteMeal(userID, id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateMeal(ctx *gin.Context) {
	id := ctx.Param("id")

	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	var body requestBody
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	updatedMeal, err := services.UpdateMeal(userID, id, body.ProductID, body.Gramms)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedMeal)
}

func Summary(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	daySummary, err := services.Summary(userID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, daySummary)
}
