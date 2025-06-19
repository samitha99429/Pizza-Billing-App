package main

import (
	"pizzabackend/database"
	"pizzabackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Enable default CORS (allow all origins, needed for React dev server)
	r.Use(cors.Default())

	// Connect to the database
	database.Connect()

	// Setup routes (your app routes)
	routes.SetupRoutes(r)

	// Run the server on port 8080
	r.Run(":3001")
}
