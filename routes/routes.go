package routes

import (
	"go-gcp/utils/random"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/random", func(c *gin.Context) {
			const numRange = 255

			c.JSON(http.StatusOK, gin.H{
				"randomNumber": random.RandomNumber(numRange),
			})
		})
	}
}
