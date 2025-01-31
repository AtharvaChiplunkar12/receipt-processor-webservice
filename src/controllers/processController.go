package controllers

import (
	"net/http"
	//"src/dtos"
	"src/models"
	
	"src/processes"

	"github.com/gin-gonic/gin"
)

// Handler for processing receipts

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := processes.GetPointsFromStore(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error":"Receipt ID not found"})
		return
	}
	c.JSON(http.StatusOK, models.PointsResponse{Points:points})
}



func ProcessReceipt(c *gin.Context) {
	// var receipt dtos.ReceiptDTO
	// if err := c.ShouldBindJSON(&receipt); err != nil {
	// 	c.JSON(400, gin.H{"error": "Invalid receipt format"})
	// 	return
	// }
	
}
