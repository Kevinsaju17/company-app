package main

import (
	"company-app/backend/config"
	"company-app/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect database
	config.ConnectDatabase()

	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Public auth routes
	routes.AuthRoutes(r)

	// Protected profile routes
	routes.ProfileRoutes(r)

	// Start server
	r.Run(":8080")
}