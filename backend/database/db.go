package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
    "pizzabackend/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=samithadilshan password=Sam123 dbname=pizzadb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    db.AutoMigrate(&models.Item{}, &models.Invoice{}, &models.InvoiceItem{})
    DB = db
}
