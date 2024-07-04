package user

import (
	"community_voice/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchUser struct {
	Username string `uri:"username" binding:"required"`
}

type ResponseNoReports struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

type Response struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Name     string          `json:"name"`
	Reports  []models.Report `json:"reports"`
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

	posts := c.DefaultQuery("posts", "false")
	if posts == "true" {
		var user models.User
		r.Find(&user, "username = ?", search.Username)
		var userReports []models.Report
		r.Find(&userReports, "username = ?", user.Username)
		c.JSON(200, Response{
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Name:     user.Name,
			Reports:  userReports,
		})
		return
	}

	var user models.User
	r.Find(&user, "username = ?", search.Username)
	if user.Username == "" {
		c.JSON(404, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(200, ResponseNoReports{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Name:     user.Name,
	})
}
