package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sandu9618/food-ordering-backend/internal/tenant"
)

func TenantMiddleware(repo *tenant.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		parts := strings.Split(host, ".")
		if len(parts) < 2 {
			c.AbortWithStatusJSON(400, gin.H{
				"error": "Invalid subdomain",
			})
			return
		}

		subdomain := parts[0]

		foundTenant, err := repo.FindBySubdomain(subdomain)

		if err != nil || foundTenant == nil {
			c.AbortWithStatusJSON(404, gin.H{
				"error": "tenant not found",
			})
			return
		}

		c.Set("tenant_id", foundTenant.ID)

		c.Next()
	}
}
