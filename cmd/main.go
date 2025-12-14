package main

import (
	"log"
	"project_calorie_tracker/internal/handlers"
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

	handlers.RegisterProoductRoutes(server)

	server.Run()
}