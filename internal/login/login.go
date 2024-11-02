package login

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
)

func getToken(username string) (string, error) {
	token, err := GenerateJwt(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (db *Struct) TryLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{"msg": "Invalid request"})
		return
	}

	var u models.User
	db.Find(&u, "username = ?", username)
	if ComparePassword(u.Password, password) {
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
