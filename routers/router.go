package routers

import (
	"github.com/gin-gonic/gin"
	"practive_service/context"
)

func SetupRouters(h *context.Application) *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api/practice-service")

	apiGroup.GET("/add-person", h.PersonController.AddPerson)

	return router
}
