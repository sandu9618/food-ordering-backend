package order

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func (h *Handler) Checkout(c *gin.Context) {
	userID := uint(c.MustGet("user_id").(float64))
	tenantID := c.MustGet("tenant_id").(uint)

	err := h.Service.CreateOrder(userID, tenantID)
	if err != nil {
		if errors.Is(err, ErrEmptyCart) {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "order created"})
}

func (h *Handler) GetOrders(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	orders, err := h.Service.OrdersForTenant(tenantID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	idParam := c.Param("id")
	orderID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid order id"})
		return
	}

	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	if err := h.Service.SetOrderStatus(uint(orderID), tenantID, req.Status); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "status updated"})
}
