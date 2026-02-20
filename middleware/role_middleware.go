package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleInterface, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		userRole := roleInterface.(string)

		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}