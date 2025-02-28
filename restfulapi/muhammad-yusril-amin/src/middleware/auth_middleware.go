package middleware

import (
	"restfull-api-lms/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMidlleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"Error":"Authorization tidak boleh kosong!"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"Error":"Authorization harus bearer token!"})
			c.Abort()
			return
		}

		userId, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"Error":"Token salah atau expired!"})
			c.Abort()
			return
		}
		
		c.Set("userId",userId)
		c.Next()
	}
}