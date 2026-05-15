package category

import "github.com/gin-gonic/gin"

type Handler struct {
	Repo *Repository
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tenantID := c.MustGet("tenant_id").(uint)

	newCategory := &Category{
		TenantID: tenantID,
		Name:     req.Name,
		Code:     req.Code,
	}

	err := h.Repo.Create(newCategory)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Category created successfully"})
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(uint)

	categories, err := h.Repo.FindAllByTenant(tenantID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, categories)
}
