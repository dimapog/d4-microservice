package ai

import (
	"net/http"

	"github.com/dimapog/jwt-microservice/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	aiRoutes := router.Group("/ai")
	{
		aiRoutes.POST("/personal-calculation", middleware.AuthMiddleware(), h.CalculatePersonalStatistics)
	}
}

// CalculatePersonalStatistics godoc
// @Summary Get AI-powered health analysis
// @Description Sends personal health data to the AI module and returns a calculated response
// @Tags ai
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body ai.PersonalCalculationRequest true "AI calculation request"
// @Success 200 {object} ai.PersonalCalculationResponse
// @Failure 400 {object} ai.ErrorResponse
// @Failure 401 {object} ai.ErrorResponse
// @Failure 500 {object} ai.ErrorResponse
// @Router /ai/personal-calculation [post]
func (h *Handler) CalculatePersonalStatistics(c *gin.Context) {
	var req PersonalCalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CalculatePersonalStatistics(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": result})
}
