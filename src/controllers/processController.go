package controllers

import (
	"net/http"
	"src/dtos"
	"src/models"

	"src/processes"

	"github.com/gin-gonic/gin"
)


func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := processes.GetPointsFromStore(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that ID."})
		return
	}
	c.JSON(http.StatusOK, models.PointsResponse{Points: points})
}

func ProcessReceipt(c *gin.Context) {
	var receipt_dto dtos.ReceiptDTO
	if err := c.ShouldBindJSON(&receipt_dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "The receipt is invalid.",
			"details": err.Error(),
		})
		return
	}

	receipt, errors := dtos.ReceiptDTOToReceipt(receipt_dto)

	if errors != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error in receipt JSON validation ":errors})
		return
	}
	id := processes.ReceiptProcessing(receipt)
	c.JSON(http.StatusOK, models.ReceiptResponse{ID: id})

}
