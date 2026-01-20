package handlers

import (
	"net/http"

	"github.com/sustatov027-max/project_calorie_tracker/internal/middlewares"
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"
	"github.com/sustatov027-max/project_calorie_tracker/internal/validation"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=user_handler.go -destination=mock/mock_services.go -package=mock

func RegisterUserRoutes(r *gin.Engine, h *UserHandler) {
	r.POST("/auth/register", h.RegisterUser)
	r.POST("/auth/login", h.LoginUser)
	r.GET("/me", middlewares.AuthMiddleware, h.GetUser)
}

var validator = validation.NewValidator()

type UserServ interface {
	RegisterUser(name string, age int, email string, password string, weight float64, height float64, gender string, activeDays int) (models.User, error)
	LoginUser(email string, password string) (string, error)
	GetUser(userID any) (models.User, error)
}

type UserHandler struct {
	service UserServ
}

func NewUserHandler(s UserServ) *UserHandler {
	return &UserHandler{service: s}
}

type RequestUserBody struct {
	Name       string  `json:"name" validate:"min=2"`
	Age        int     `json:"age" validate:"gte=1,lte=120"`
	Email      string  `json:"email" validate:"email"`
	Password   string  `json:"password" validate:"min=8"`
	Weight     float64 `json:"weight" validate:"gt=0"`
	Height     float64 `json:"height" validate:"gt=0"`
	Gender     string  `json:"gender"`
	ActiveDays int     `json:"activeDays" validate:"gte=0,lte=7"`
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var body RequestUserBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	validationErrors := validator.Validate(&body)
	if len(validationErrors) != 0 {
		ctx.IndentedJSON(http.StatusBadRequest, validationErrors)
		return
	}

	newUser, err := h.service.RegisterUser(body.Name, body.Age, body.Email, body.Password, body.Weight, body.Height, body.Gender, body.ActiveDays)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error register user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	type LoginRequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body LoginRequestBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	token, err := h.service.LoginUser(body.Email, body.Password)
	if err != nil {
		ctx.IndentedJSON(401, map[string]string{"Error login user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, map[string]string{"token": token, "token_type": "Bearer"})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	user, err := h.service.GetUser(userID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error get user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}
