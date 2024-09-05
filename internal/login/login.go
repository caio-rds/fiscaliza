package login

import (
	"fiscaliza/internal/models"
	"fiscaliza/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getToken(username string) (string, error) {
	token, err := services.GenerateJwt(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

type Login struct {
	*gorm.DB
}

func NewLogin(db *gorm.DB) *Login {
	value := Login{
		db,
	}
	return &value
}

func (l *Login) TryLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{"msg": "Invalid request"})
		return
	}

	var u models.User
	l.Find(&u, "username = ?", username)
	if services.ComparePassword(u.Password, password) {
		if token, err := getToken(username); err != nil {
			c.JSON(500, gin.H{"msg": "Internal server error"})
			return
		} else {
			c.JSON(200, gin.H{"token": token})
			return
		}
	}
	c.JSON(401, gin.H{"msg": "Invalid credentials"})
	return
}
