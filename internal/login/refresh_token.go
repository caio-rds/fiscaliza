package login

import (
	"errors"
	"fiscaliza/internal/auth"
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (db *Struct) RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := auth.ValidateToken(tokenString)

	var user *models.User
	db.Find(&user, "username = ?", claims.Username)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if err != nil && !errors.Is(err, auth.ErrTokenExpired) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	newToken, err := auth.GenerateJwt(claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
