package main

import (
	"deals/db"
	"deals/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db.InitDB()

	// Set up routes with controllers
	r := gin.Default()

	// Set up routes
	routes.InitRoutes(r)

	r.Run(":8086")
}
