package main

import (
	"log"
	"project_calorie_tracker/internal/handlers"
	"project_calorie_tracker/internal/repositories"
	"project_calorie_tracker/internal/services"
	"project_calorie_tracker/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.Init()
	server := gin.Default()

	userRepo := repositories.UserRepository{}
	userService := services.NewUserService(&userRepo)
	userHandler := handlers.NewUserHandler(userService)

	handlers.RegisterProductRoutes(server)
	handlers.RegisterUserRoutes(server, userHandler)
	handlers.RegisterDiaryRoutes(server)

	server.Run()
}