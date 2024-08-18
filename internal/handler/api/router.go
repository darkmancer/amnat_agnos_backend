package api

import (
	"strong_password_recommendation/internal/core/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, passwordService *service.PasswordService) {
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		HandleStrongPasswordSteps(c, passwordService)
	})
}
