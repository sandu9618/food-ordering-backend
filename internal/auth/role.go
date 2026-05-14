package auth

import "github.com/gin-gonic/gin"

func RoleMiddleware(
	allowedRoles ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)

		for _, allowed := range allowedRoles {
			if allowed == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
	}
}
