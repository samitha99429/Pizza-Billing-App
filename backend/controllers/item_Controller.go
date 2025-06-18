package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "pizzabackend/models"
    "pizzabackend/database"
)

func GetItems(c *gin.Context) {
    var items []models.Item
    database.DB.Find(&items)
    c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&item)
    c.JSON(http.StatusOK, item)
}
