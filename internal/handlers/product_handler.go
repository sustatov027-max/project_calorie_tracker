package handlers

import (
	"net/http"
	"project_calorie_tracker/internal/middlewares"
	"project_calorie_tracker/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.POST("/products", middlewares.AuthMiddleware, CreateProduct)
	r.GET("/products", middlewares.AuthMiddleware, GetAllProducts)
	r.DELETE("/products/:id", middlewares.AuthMiddleware, DeleteProduct)
	r.PUT("/products/:id", middlewares.AuthMiddleware, UpdateProduct)
}

type RequestProductBody struct {
	Name          string  `json:"name"`
	Calories      float64 `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func CreateProduct(ctx *gin.Context) {
	var body RequestProductBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	newProduct, err := services.CreateProduct(body.Name, body.Calories, body.Proteins, body.Fats, body.Carbohydrates)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error create product": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newProduct)
}

func GetAllProducts(ctx *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error read product table": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, products)
}

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := services.DeleteProduct(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error delete product": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateProduct(ctx *gin.Context) {
	var body RequestProductBody
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read id": err.Error()})
		return
	}

	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	updateProduct, err := services.UpdateProduct(id, body.Name, body.Calories, body.Proteins, body.Fats, body.Carbohydrates)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error update product": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updateProduct)
}
