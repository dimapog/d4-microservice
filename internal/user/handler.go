package user

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
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/", h.CreateUser)
		userRoutes.GET("/:id", middleware.AuthMiddleware(), h.GetUserByID)
		userRoutes.PATCH("/", middleware.AuthMiddleware(), h.PatchUser)
		userRoutes.DELETE("/:id", h.DeleteUser)
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Register a new user with name, email, and password
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.CreateUserRequest true "Create user request"
// @Success 201 {object} user.UserResponse
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /user/ [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieve a user profile by ID
// @Tags user
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.UserResponse
// @Failure 401 {object} user.ErrorResponse
// @Failure 404 {object} user.ErrorResponse
// @Router /user/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 400 {object} user.ErrorResponse
// @Failure 500 {object} user.ErrorResponse
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}

// PatchUser godoc
// @Summary Update authenticated user
// @Description Update fields for the authenticated user
// @Tags user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body user.UpdateUserRequest true "Update user request"
// @Success 200 {object} user.UserResponse
// @Failure 400 {object} user.ErrorResponse
// @Failure 401 {object} user.ErrorResponse
// @Failure 404 {object} user.ErrorResponse
// @Router /user/ [patch]
func (h *Handler) PatchUser(c *gin.Context) {
	// Get authenticated user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID is already uint from middleware
	authUserID := userID.(uint)

	user, err := h.service.UpdateUser(authUserID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
