package middleware

import (
	"backend/internal/entity"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userAny, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		user, ok := userAny.(*entity.User)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if user.Role != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "Not enough permissions"})
			return
		}

		c.Next()
	}
}
