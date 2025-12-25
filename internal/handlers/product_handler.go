package handlers

import (
	"net/http"
	"project_calorie_tracker/internal/middlewares"
	"project_calorie_tracker/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, h *ProductHandler) {
	r.POST("/products", middlewares.AuthMiddleware, h.CreateProduct)
	r.GET("/products", middlewares.AuthMiddleware, h.GetAllProducts)
	r.DELETE("/products/:id", middlewares.AuthMiddleware, h.DeleteProduct)
	r.PUT("/products/:id", middlewares.AuthMiddleware, h.UpdateProduct)
}

type ProductServ interface{
	CreateProduct(name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error)
	GetAllProducts() ([]models.Product, error)
	DeleteProduct(id string) error
	UpdateProduct(id int, name string, calories float64, proteins float64, fats float64, carbohydrates float64) (models.Product, error)
	GetProductByID(id int) (models.Product, error)
}

type ProductHandler struct{
	service ProductServ
}

func NewProductHandler(s ProductServ) *ProductHandler{
	return &ProductHandler{service: s}
}

type RequestProductBody struct {
	Name          string  `json:"name"`
	Calories      float64 `json:"calories"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var body RequestProductBody

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, map[string]string{"Error read request body": err.Error()})
		return
	}

	newProduct, err := h.service.CreateProduct(body.Name, body.Calories, body.Proteins, body.Fats, body.Carbohydrates)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error create product": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newProduct)
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error read product table": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, products)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.DeleteProduct(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error delete product": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
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

	updateProduct, err := h.service.UpdateProduct(id, body.Name, body.Calories, body.Proteins, body.Fats, body.Carbohydrates)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, map[string]string{"Error update product": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updateProduct)
}
