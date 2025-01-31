package dtos

type ReceiptDTO struct {
    Retailer     string    `json:"retailer" validate:"required"`
    PurchaseDate string    `json:"purchaseDate" validate:"required,datetime=2022-01-01"`
    PurchaseTime string    `json:"purchaseTime" validate:"required,datetime=13:01"`
    Items        []ItemDTO `json:"items" validate:"required,dive"`
	Total        float64   `json:"total" validate:"required"`
}

