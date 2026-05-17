package cart

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *Repository
}

type AddItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (h *Handler) AddItem(c *gin.Context) {

	var req AddItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	userID := uint(c.MustGet("user_id").(float64))
	tenantID := c.MustGet("tenant_id").(uint)

	cart, err := h.Repo.FindOrCreateCart(
		userID,
		tenantID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	item := &CartItem{
		CartID:    cart.ID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err = h.Repo.AddItem(item)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(201, gin.H{
		"message": "item added",
	})
}

func (h *Handler) GetCart(c *gin.Context) {

	userID := uint(c.MustGet("user_id").(float64))
	tenantID := c.MustGet("tenant_id").(uint)

	cart, err := h.Repo.GetCart(
		userID,
		tenantID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, cart)
}
