package main

import (
	"github.com/andito28/RestAPI_Golang/controllers"
	"github.com/andito28/RestAPI_Golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.SetupDB()
	// db.AutoMigrate(models.Product{})

	productController := controllers.NewProductController(db)
	router := gin.Default()
	api := router.Group("api/v1")
	api.GET("/test", productController.Index)

	router.Run("localhost:8080")
}
