package main

import (
	"example.com/rest-api-go/db"
	"example.com/rest-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
