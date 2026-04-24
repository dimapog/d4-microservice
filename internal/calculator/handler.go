package calculator

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
	calcRoutes := router.Group("/calculator")
	{
		calcRoutes.GET("/bmi", middleware.AuthMiddleware(), h.CalculateBMI)
		calcRoutes.GET("/hrz", middleware.AuthMiddleware(), h.CalculateHeartRateZones)
	}
}

// CalculateBMI godoc
// @Summary Calculate BMI for authenticated user
// @Description Computes body mass index from the authenticated user's stored profile
// @Tags calculator
// @Security BearerAuth
// @Produce json
// @Success 200 {object} calculator.BMIResponse
// @Failure 401 {object} calculator.ErrorResponse
// @Failure 400 {object} calculator.ErrorResponse
// @Router /calculator/bmi [get]
func (h *Handler) CalculateBMI(c *gin.Context) {
	// Get authenticated user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	authUserID := userID.(uint)

	resp, err := h.service.CalculateBMIByUserID(authUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CalculateHeartRateZones godoc
// @Summary Calculate heart rate zones for authenticated user
// @Description Computes heart rate training zones from the authenticated user's stored profile
// @Tags calculator
// @Security BearerAuth
// @Produce json
// @Success 200 {object} calculator.HRZResponse
// @Failure 401 {object} calculator.ErrorResponse
// @Failure 400 {object} calculator.ErrorResponse
// @Router /calculator/hrz [get]
func (h *Handler) CalculateHeartRateZones(c *gin.Context) {
	// Get authenticated user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	authUserID := userID.(uint)

	resp, err := h.service.CalculateHeartRateZonesByUserID(authUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
