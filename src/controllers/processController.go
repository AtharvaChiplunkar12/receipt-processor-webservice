package controllers

import (
	"net/http"
	"regexp"
	"src/dtos"
	"src/models"
	"time"

	"src/processes"

	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := processes.GetPointsFromStore(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID"})
		return
	}
	c.JSON(http.StatusOK, models.PointsResponse{Points: points})
}

func ProcessReceipt(c *gin.Context) {
	var receipt_dto dtos.ReceiptDTO
	if err := c.ShouldBindJSON(&receipt_dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"The receipt is invalid": err.Error(),
		})
		return
	}

	if validationErrors := validateReceipt(receipt_dto); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"The receipt is invalid": validationErrors,
		})
		return
	}

	receipt, errors := dtos.ReceiptDTOToReceipt(receipt_dto)

	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"The receipt is invalid": errors})
		return
	}
	id := processes.ReceiptProcessing(receipt)
	c.JSON(http.StatusOK, models.ReceiptResponse{ID: id})

}

func validateReceipt(receipt dtos.ReceiptDTO) []string {
	var errors []string

	// Validate retailer name pattern
	retailerPattern := regexp.MustCompile(`^[\w\s\-&]+$`)
	if receipt.Retailer == "" {
		errors = append(errors, "Retailer name required")
	} else if !retailerPattern.MatchString(receipt.Retailer) {
		errors = append(errors, "Invalid retailer name format.")
	}

	// Validate purchaseDate format (YYYY-MM-DD)
	if receipt.PurchaseDate == "" {
		errors = append(errors, "Purchase Date required")
	} else if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		errors = append(errors, "Invalid purchaseDate format. Expected YYYY-MM-DD.")
	}

	// Validate purchaseTime format (HH:MM in 24-hour format)
	if receipt.PurchaseTime == "" {
		errors = append(errors, "Purchase Time required")
	}else if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		errors = append(errors, "Invalid purchaseTime format. Expected HH:MM in 24-hour format.")
	}

	// Validate total amount pattern (digits with two decimal places)
	totalPattern := regexp.MustCompile(`^\d+\.\d{2}$`)
	if receipt.Total == "" {
		errors = append(errors, "Total required")
	} else if !totalPattern.MatchString(receipt.Total) {
		errors = append(errors, "Invalid total format. Expected numeric value with two decimal places.")
	}

	// Validate at least one item exists
	if len(receipt.Items) == 0 {
		errors = append(errors, "Receipt must contain at least one item.")
	}

	return errors
}
