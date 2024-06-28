package user

import (
	"community_voice/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchUser struct {
	Username string `uri:"username" binding:"required"`
}

type read struct {
	*gorm.DB
}

func NewRead(db *gorm.DB) *read {
	value := read{
		db,
	}
	return &value
}

func (r *read) Read(c *gin.Context) {
	var search SearchUser
	if err := c.ShouldBindUri(&search); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var user models.User
	r.Find(&user, "username = ?", search.Username)

	c.JSON(200, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.Phone,
		"name":     user.Name,
	})
}
