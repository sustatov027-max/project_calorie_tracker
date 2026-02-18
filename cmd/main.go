package main

import (
	"log"

	"github.com/sustatov027-max/project_calorie_tracker/internal/config"
	"github.com/sustatov027-max/project_calorie_tracker/internal/handlers"
	"github.com/sustatov027-max/project_calorie_tracker/internal/repositories"
	"github.com/sustatov027-max/project_calorie_tracker/internal/services"
	"github.com/sustatov027-max/project_calorie_tracker/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	database.Init()
	server := gin.New()
	server.Use(gin.Logger())

	userRepo := repositories.UserRepository{}
	userService := services.NewUserService(&userRepo)
	userHandler := handlers.NewUserHandler(userService)

	productRepo := repositories.ProductRepository{}
	productService := services.NewProductService(&productRepo)
	productHandler := handlers.NewProductHandler(productService)

	diaryRepo := repositories.DiaryRepository{}
	diaryService := services.NewDiaryService(&diaryRepo)
	diaryHandler := handlers.NewDiaryHandler(diaryService)

	handlers.RegisterProductRoutes(server, productHandler)
	handlers.RegisterUserRoutes(server, userHandler)
	handlers.RegisterDiaryRoutes(server, diaryHandler)

	if err := server.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server startup error: %v", err)
	}
}
