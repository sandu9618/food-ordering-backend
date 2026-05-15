package product

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	Service *Service
}

type CreateProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
}

func (h *Handler) Createproduct(c *gin.Context) {
	var req CreateProductReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	newProduct := &Product{
		TenantID:    tenantID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	err := h.Service.Create(newProduct)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Product created successfully"})
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	products, err := h.Service.FindAllByTenantId(tenantID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

func (h *Handler) GetProductById(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	tenantID := c.MustGet("tenant_id").(uint)

	product, err := h.Service.FindById(uint(id), tenantID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	existing, err := h.Service.FindById(uint(id), tenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "product not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var req CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existing.Name = req.Name
	existing.Description = req.Description
	existing.Price = req.Price
	existing.CategoryID = req.CategoryID

	if err := h.Service.Update(existing); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	err = h.Service.Delete(uint(id), tenantID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
