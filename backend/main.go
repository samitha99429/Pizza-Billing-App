package main

import (
	"pizzabackend/database"
     "pizzabackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r:=gin.Default()
	database.Connect()
	routes.SetupRoutes(r)
    r.Run(":8080")
}