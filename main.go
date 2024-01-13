package main

import (
	"github.com/gin-gonic/gin"
	"wallnut-studios.com/rest-api/db"
	"wallnut-studios.com/rest-api/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8000")
}
