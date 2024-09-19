package user

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
)

func (db *Struct) Restore(c *gin.Context) {
	userNameParam := c.Param("user")
	if userNameParam == "" {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := db.DB.Unscoped().Where("username = ?", userNameParam).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	if user.DeletedAt == nil {
		c.JSON(400, gin.H{"error": "User not deleted"})
		return
	}

	user.DeletedAt = nil
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User restored"})
}
