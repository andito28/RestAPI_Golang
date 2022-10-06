package controllers

import (
	"net/http"

	"github.com/andito28/RestAPI_Golang/helper"
	"github.com/andito28/RestAPI_Golang/models"
	"github.com/gin-gonic/gin"
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

func FormatProduct(product models.Product) ProductFormatter {
	var formatter ProductFormatter
	formatter.ID = product.ID
	formatter.ProductName = product.ProductName
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

func (ctx *productController) Index(c *gin.Context) {
	var products []models.Product
	ctx.db.Find(&products)
	response := helper.ApiResponse("List Product", http.StatusOK, "Success", FormatProducts(products))
	c.JSON(http.StatusOK, response)
}
