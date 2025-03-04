package middleware

import (
	"resfulapi/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMidlleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthHeader := c.GetHeader("Authorization")

		if AuthHeader == "" {
			c.JSON(401, gin.H{"Error": "Authorization tidak boleh kosong!"})
			c.Abort()
			return
		}

		Parts := strings.Split(AuthHeader, " ")
		if len(Parts) != 2 || Parts[0] != "Bearer" {
			c.JSON(401, gin.H{"Error": "Authorization harus bearer token!"})
			c.Abort()
			return
		}

		UserId, err := utils.ValidateToken(Parts[1])
		if err != nil {
			c.JSON(401, gin.H{"Error": "Token salah atau expired!"})
			c.Abort()
			return
		}

		c.Set("UserId", UserId)
		c.Next()
	}
}
