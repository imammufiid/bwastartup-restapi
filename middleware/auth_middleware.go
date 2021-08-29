package middleware

import (
	"bwastartup/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// check if authorization contains "Bearer" text
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// split Bearer token
	token := splitBearerToken(authHeader)
	if len(strings.Trim(token, " ")) <= 0 {
		response := helper.ApiResponse("Invalid Token Parsing", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
}

func splitBearerToken(bearerToken string) (string) {
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}

	return ""
}