package handlers

import (
	"github.com/sustatov027-max/project_calorie_tracker/internal/middlewares"
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"
	"github.com/sustatov027-max/project_calorie_tracker/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterDiaryRoutes(r *gin.Engine, h *DiaryHandler) {
	r.POST("/diary", middlewares.AuthMiddleware, h.CreateMeal)
	r.GET("/diary", middlewares.AuthMiddleware, h.GetMealsForDay)
	r.DELETE("/diary/:id", middlewares.AuthMiddleware, h.DeleteMeal)
	r.PUT("/diary/:id", middlewares.AuthMiddleware, h.UpdateMeal)
	r.GET("/diary/summary", middlewares.AuthMiddleware, h.Summary)
}

type DiaryServ interface {
	CreateMeal(userID int, productID int, gramms float64) (models.MealLog, error)
	GetAllMealsForDay(userID int, date time.Time) ([]models.MealLog, error)
	DeleteMeal(userID int, id string) error
	UpdateMeal(userID int, id string, productID int, gramms float64) (models.MealLog, error)
	Summary(userID int) (models.DaySummary, error)
}

type DiaryHandler struct {
	service DiaryServ
}

func NewDiaryHandler(s DiaryServ) *DiaryHandler {
	return &DiaryHandler{service: s}
}

type requestBody struct {
	ProductID int     `json:"product_id"`
	Gramms    float64 `json:"gramms"`
}

func (h *DiaryHandler) CreateMeal(ctx *gin.Context) {
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

	createdMeal, err := h.service.CreateMeal(userID, body.ProductID, body.Gramms)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdMeal)
}

func (h *DiaryHandler) GetMealsForDay(ctx *gin.Context) {
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

	meals, err := h.service.GetAllMealsForDay(userID, date)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, meals)
}

func (h *DiaryHandler) DeleteMeal(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	err = h.service.DeleteMeal(userID, id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *DiaryHandler) UpdateMeal(ctx *gin.Context) {
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

	updatedMeal, err := h.service.UpdateMeal(userID, id, body.ProductID, body.Gramms)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedMeal)
}

func (h *DiaryHandler) Summary(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	daySummary, err := h.service.Summary(userID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, daySummary)
}
