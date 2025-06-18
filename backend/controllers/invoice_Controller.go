package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "pizzabackend/models"
    "pizzabackend/database"
)

// func CreateInvoice(c *gin.Context) {
//     var invoice models.Invoice
//     if err := c.ShouldBindJSON(&invoice); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     database.DB.Create(&invoice)
//     c.JSON(http.StatusOK, invoice)
// }


func CreateInvoice(c *gin.Context) {
    var invoice models.Invoice
    if err := c.ShouldBindJSON(&invoice); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the invoice first
    if err := database.DB.Create(&invoice).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Now assign the invoice ID to each invoice item and save them
    for i := range invoice.Items {
        invoice.Items[i].InvoiceID = invoice.ID
        if err := database.DB.Create(&invoice.Items[i]).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, invoice)
}

func GetInvoices(c *gin.Context) {
    var invoices []models.Invoice
    database.DB.Preload("Items.Item").Find(&invoices)
    c.JSON(http.StatusOK, invoices)
}
