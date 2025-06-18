package models

import "time"

type Invoice struct {
    ID           uint         `json:"id" gorm:"primaryKey"`
    CustomerName string       `json:"customer_name"`
    Tax          float64      `json:"tax"`
    TotalAmount  float64      `json:"total_amount"`
    CreatedAt    time.Time    `json:"created_at"`
    Items        []InvoiceItem `json:"items"`
}

type InvoiceItem struct {
    ID        uint `json:"id" gorm:"primaryKey"`
    InvoiceID uint
    ItemID    uint `json:"item_id"` 
    Quantity  int
    Item      Item `json:"item" gorm:"foreignKey:ItemID"`
}
