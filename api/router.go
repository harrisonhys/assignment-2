package api

import (
	"assignment-two/controller/general"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/api/data", general.GetData)
	router.GET("/api/data/:id", general.ShowData)
	router.POST("/api/data", general.CreateData)
	router.PUT("/api/data/:id", general.UpdateData)
	router.DELETE("/api/data/:id", general.DeleteData)

	return router
}
