package routes

import (
    "github.com/gin-gonic/gin"
	"pizzabackend/controllers"

)

func SetupRoutes(r *gin.Engine) {
    r.GET("/items", controllers.GetItems)
    r.POST("/items", controllers.CreateItem)
    r.POST("/invoices", controllers.CreateInvoice)
    r.GET("/invoices", controllers.GetInvoices)
}
