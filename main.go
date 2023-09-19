package main

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-assignment2/app/configs"
)

func main() {
	configs.GetDatabaseConnection()
	r := gin.Default()

	r.Run()
}
