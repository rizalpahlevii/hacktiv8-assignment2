package main

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/config"
)

func main() {
	config.GetDatabaseConnection()
	r := gin.Default()

	r.Run()
}
