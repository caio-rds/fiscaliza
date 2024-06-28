package login

import (
	"community_voice/internal/auth"
	"community_voice/internal/hash"
	"community_voice/internal/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getToken(username string) (string, error) {
	token, err := auth.GenerateJwt(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

type login struct {
	*gorm.DB
}

func NewLogin(db *gorm.DB) *login {
	value := login{
		db,
	}
	return &value
}

func (l *login) TryLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{"msg": "Invalid request"})
		return
	}

	var user user.User
	l.Find(&user, "username = ?", username)
	fmt.Println(hash.ComparePassword(user.Password, password))
	if hash.ComparePassword(user.Password, password) {
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
