package api

import (
	"net/http"
	"strong_password_recommendation/internal/core/service"
	"strong_password_recommendation/internal/handler/dto"

	"github.com/gin-gonic/gin"
)

func HandleStrongPasswordSteps(c *gin.Context, service *service.PasswordService) {
	var req dto.PasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	steps := service.CalculateSteps(req.InitPassword)
	c.JSON(http.StatusOK, dto.PasswordResponse{NumOfSteps: steps})
}
