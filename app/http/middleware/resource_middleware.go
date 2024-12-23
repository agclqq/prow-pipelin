package middleware

import (
	"github.com/gin-gonic/gin"
	"your_project/internal/domain/repositories"
	"your_project/internal/domain/services"
)

func ResourceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		stageRepo := repositories.NewInMemoryConfFlowStageRepository()
		stageService := services.NewConfFlowStageService(stageRepo)
		c.Set("stageService", stageService)
		c.Next()
	}
} 