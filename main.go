package main

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/app/models"
	"hacktiv8-assignment2/config"
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

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
