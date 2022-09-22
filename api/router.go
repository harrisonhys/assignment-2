package api

import (
	"github.com/gin-gonic/gin"
	"assignment-two/controller/general"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/api/data", general.GetData)
	// router.GET("/api/data/:id", general.showData)
	// router.POST("/api/data", general.createData)
	// router.PUT("/api/data/:id", general.UpdateData)
	// router.DELETE("/api/data", general.deleteData)

	return router
}