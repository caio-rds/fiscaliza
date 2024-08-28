package login

import (
	"errors"
	"fiscaliza/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := services.ValidateToken(tokenString)
	if err != nil && !errors.Is(err, services.ErrTokenExpired) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	newToken, err := services.GenerateJwt(claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
