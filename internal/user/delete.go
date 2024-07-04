package user

import (
	"community_voice/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type uDelete struct {
	*gorm.DB
}

type DeleteUser struct {
	Username string `json:"username" uri:"username" binding:"required"`
}

func NewDelete(db *gorm.DB) *uDelete {
	value := uDelete{
		db,
	}
	return &value
}

func (d *uDelete) Delete(c *gin.Context, username string) {
	var deleteUser DeleteUser
	if err := c.ShouldBindUri(&deleteUser); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if deleteUser.Username == "" {
		c.JSON(400, gin.H{"msg": "username is required"})
		return
	}

	if deleteUser.Username != username {
		c.JSON(400, gin.H{"msg": "username in url and body should be the same"})
		return
	}

	var user *models.User
	if err := d.Find(&user, "username = ?", deleteUser.Username); err.Error != nil {
		c.JSON(400, gin.H{"msg": "user not found"})
		return
	}
	user.DeletedAt = &gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}
	if err := d.Save(&user); err.Error != nil {
		c.JSON(500, gin.H{"msg": err.Error})
		return
	}
	c.JSON(200, gin.H{"msg": "user deleted"})
}
