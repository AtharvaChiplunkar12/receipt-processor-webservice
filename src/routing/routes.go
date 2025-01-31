package routing

import (
	"src/controllers"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", controllers.ProcessReceipt)
	router.GET("/receipts/{id}/points", controllers.GetPoints)
	return router
}
