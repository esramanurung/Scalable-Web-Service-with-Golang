package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartingServer() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrder)
	router.GET("/order/:orderID", controllers.GetOrderByID)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
