package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hacktiv8-assignment2/config"
	"hacktiv8-assignment2/docs"
	"hacktiv8-assignment2/models"
	"hacktiv8-assignment2/routes"
)

func main() {
	db := config.GetDatabaseConnection()

	err := db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	routes.OrderRoutes(router)

	docs.SwaggerInfo.Title = "Hacktiv8 Assignment 2"
	docs.SwaggerInfo.Description = "Order API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Order API",
		})
	})
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
