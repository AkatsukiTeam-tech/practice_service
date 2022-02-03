package middleware

import (
	"github.com/gin-gonic/gin"
	"practive_service/utils"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		ValidToken(context)
		context.Next()
	}
}

func ValidToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	validToken := "Basic " + utils.GetEnv("AUTH_TOKEN", "dGF5eWFiOnRheXlhYjEyMyQjQA==")

	if token == validToken {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}
