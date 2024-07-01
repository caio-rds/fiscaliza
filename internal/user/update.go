package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
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

func (u *update) UpdateUser(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}
	if err := u.DB.Model(&User{}).Where("username = ?", req.Username).
		Updates(User{
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
