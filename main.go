package main

import (
	"goapi/db"
	"goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisteredRoutes(server)

	server.Run(":8080")
}
