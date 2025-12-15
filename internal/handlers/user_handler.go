package handlers

import (
	"net/http"
	"project_calorie_tracker/internal/middlewares"
	"project_calorie_tracker/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/auth/register", RegisterUser)
	r.POST("/auth/login", LoginUser)
	r.GET("/me", middlewares.AuthMiddleware, GetUser)
}

type RequestUserBody struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Weight     float64 `json:"weight"`
	Height     float64 `json:"height"`
	Gender     string  `json:"gender"`
	ActiveDays int     `json:"activeDays"`
}

func RegisterUser(ctx *gin.Context) {
	var body RequestUserBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	newUser, err := services.RegisterUser(body.Name, body.Age, body.Email, body.Password, body.Weight, body.Height, body.Gender, body.ActiveDays)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error register user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func LoginUser(ctx *gin.Context) {
	type LoginRequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body LoginRequestBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	token, err := services.LoginUser(body.Email, body.Password)
	if err != nil{
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error login user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, map[string]string{"token": token, "token_type":"Bearer"})
}

func GetUser(ctx *gin.Context){
	userID, _ := ctx.Get("userID")

	user, err := services.GetUser(userID)
	if err != nil{
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error get user": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}
