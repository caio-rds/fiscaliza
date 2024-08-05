package user

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UpdateRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type update struct {
	*gorm.DB
}

func Update(db *gorm.DB) *update {
	value := update{
		db,
	}
	return &value
}

func (db *Struct) UpdateUser(c *gin.Context, username string) {
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}
	if err := db.DB.Model(&models.User{}).Where("username = ?", username).
		Updates(models.User{
			Email: req.Email,
			Name:  req.Name,
			Phone: req.Phone,
		}).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User updated",
	})
	return
}
