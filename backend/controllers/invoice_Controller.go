package controllers

import (
    "bytes"
    "io"
	"log"
	"net/http"
	"pizzabackend/database"
	"pizzabackend/models"

	"github.com/gin-gonic/gin"
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


// func CreateInvoice(c *gin.Context) {
//     var invoice models.Invoice
//     if err := c.ShouldBindJSON(&invoice); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

//         log.Println("BindJSON Error:", err)
//         return
//     }


//     // Save the invoice first
//     if err := database.DB.Create(&invoice).Error; err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Now assign the invoice ID to each invoice item and save them
//     for i := range invoice.Items {
//         invoice.Items[i].InvoiceID = invoice.ID
//         if err := database.DB.Create(&invoice.Items[i]).Error; err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//     }

//     c.JSON(http.StatusOK, invoice)
// }
// func CreateInvoice(c *gin.Context) {
//     // Step 1: Log raw JSON input
//     var buf bytes.Buffer
//     tee := io.TeeReader(c.Request.Body, &buf)
//     bodyBytes, _ := io.ReadAll(tee)
//     log.Println("RAW JSON INPUT:", string(bodyBytes))

//     // Step 2: Restore the body for binding
//     c.Request.Body = io.NopCloser(&buf)

//     var invoice models.Invoice

//     // Step 3: Bind JSON
//     if err := c.ShouldBindJSON(&invoice); err != nil {
//         log.Println("BindJSON Error:", err)
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // Step 4: Save invoice
//     if err := database.DB.Create(&invoice).Error; err != nil {
//         log.Println("Invoice Save Error:", err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Step 5: Save each item
//     for i := range invoice.Items {
//         invoice.Items[i].InvoiceID = invoice.ID
//         if err := database.DB.Create(&invoice.Items[i]).Error; err != nil {
//             log.Println("InvoiceItem Save Error:", err)
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//     }

//     // Step 6: Re-fetch full invoice with nested items
//     var fullInvoice models.Invoice
//     if err := database.DB.Preload("Items.Item").First(&fullInvoice, invoice.ID).Error; err != nil {
//         log.Println("Preload Error:", err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Step 7: Respond with complete invoice
//     c.JSON(http.StatusOK, fullInvoice)
// }


func CreateInvoice(c *gin.Context) {
    var invoice models.Invoice

    // Log raw JSON
    body, _ := io.ReadAll(c.Request.Body)
    log.Println("RAW JSON INPUT:", string(body))
    c.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // restore body

    // Bind JSON
    if err := c.ShouldBindJSON(&invoice); err != nil {
        log.Println("BindJSON Error:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save invoice
    if err := database.DB.Create(&invoice).Error; err != nil {
        log.Println("Invoice Save Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Save items safely
    var invoiceItems []models.InvoiceItem
    for _, item := range invoice.Items {
        invoiceItems = append(invoiceItems, models.InvoiceItem{
            InvoiceID: invoice.ID,
            ItemID:    item.ItemID,
            Quantity:  item.Quantity,
        })
    }

    if err := database.DB.Create(&invoiceItems).Error; err != nil {
        log.Println("InvoiceItem Bulk Save Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Fetch back full invoice with items
    var fullInvoice models.Invoice
    if err := database.DB.Preload("Items.Item").First(&fullInvoice, invoice.ID).Error; err != nil {
        log.Println("Preload Error:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, fullInvoice)
}

func GetInvoices(c *gin.Context) {
    var invoices []models.Invoice
    database.DB.Preload("Items.Item").Find(&invoices)
    c.JSON(http.StatusOK, invoices)
}


