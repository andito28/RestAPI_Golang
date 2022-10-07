package controllers

import (
	"net/http"

	"github.com/andito28/RestAPI_Golang/helper"
	"github.com/andito28/RestAPI_Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type productController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *productController {
	return &productController{db}
}

type ProductFormatter struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type ProductInput struct {
	ProductName string `json:"product_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

func FormatProduct(product models.Product) ProductFormatter {
	var formatter ProductFormatter
	formatter.ID = product.ID
	formatter.ProductName = product.ProductName
	formatter.Price = product.Price
	formatter.Quantity = product.Quantity
	return formatter
}

func FormatProducts(products []models.Product) []ProductFormatter {
	if len(products) == 0 {
		return []ProductFormatter{}
	}
	var productsFormatter []ProductFormatter
	for _, product := range products {
		formatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, formatter)
	}
	return productsFormatter
}

func (ctr *productController) Index(c *gin.Context) {
	var products []models.Product
	ctr.db.Find(&products)
	response := helper.ApiResponse("List Product", http.StatusOK, "Success", FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (ctr *productController) Store(c *gin.Context) {
	var input ProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		response := helper.ApiResponse("List Product", http.StatusOK, "Success", errors)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
