package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sandu9618/food-ordering-backend/internal/user"
)

type Handler struct {
	Service *Service
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	newUser := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		TenantID: tenantID,
		Role:     req.Role,
	}

	err := h.Service.Register(newUser)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})

}
