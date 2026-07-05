package routes

import (
	"go-gcp/config"
	"go-gcp/utils/auth"
	"go-gcp/utils/random"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/login", func(c *gin.Context) {
			token, err := auth.GenerateToken(1, cfg.JWTSecret)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": token})
		})

		protected := api.Group("/")
		protected.Use(AuthMiddleware(cfg))
		{
			protected.GET("/random", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"randomNumber": random.RandomNumber(255),
				})
			})
		}
	}
}
