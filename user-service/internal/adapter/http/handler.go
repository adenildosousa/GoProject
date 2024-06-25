package http

import (
	"net/http"
	"user-service/internal/core/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(r *gin.Engine, us service.UserService) {
	h := &Handler{userService: us}
	r.GET("/profile", h.Profile)
}

func (h *Handler) Profile(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	user, err := h.userService.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username})
}
