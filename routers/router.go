package routers

import (
	"github.com/gin-gonic/gin"
	"practive_service/context"
	"practive_service/middleware"
)

func SetupRouters(h *context.Application) *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api/practice-service")
	apiGroup.Use(middleware.TokenAuthMiddleware())

	apiGroup.POST("/add-person", h.PersonController.AddPerson)
	apiGroup.GET("/get-person", h.PersonController.FindPersonById)
	apiGroup.PUT("/update-person", h.PersonController.UpdatePerson)
	apiGroup.DELETE("/delete-person", h.PersonController.DeletePerson)
	apiGroup.GET("/get-all-persons", h.PersonController.GetAllPerson)
	apiGroup.GET("/get-all-persons-query", h.PersonController.GetAllPersonByQuery)

	return router
}
