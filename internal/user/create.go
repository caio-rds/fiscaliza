package user

import (
	"community_voice/internal/hash"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/mail"
	"regexp"
	"strings"
)

type CreateReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

func (u *CreateReq) validate() (bool, error) {

	specialChars := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};:'",<>./?\\|]`)
	upperCase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false, fmt.Errorf("invalid email")
	}
	if len(u.Username) < 3 {
		return false, fmt.Errorf("username must have at least 3 characters")
	}
	if len(u.Password) < 6 {
		return false, fmt.Errorf("password must have at least 6 characters")
	}

	if strings.Contains(u.Password, "password") {
		return false, fmt.Errorf("password cannot contain the word 'password'")
	}

	if !specialChars.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one special character")
	}

	if !number.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one number")
	}

	if !upperCase.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one uppercase letter")
	}

	return true, nil
}

type create struct {
	*gorm.DB
}

func NewCreate(db *gorm.DB) *create {
	value := create{
		db,
	}
	return &value
}

func (cr *create) Create(c *gin.Context) {
	var data CreateReq
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}
	if _, err := data.validate(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := cr.DB.Where("username = ?", data.Username).Or("email = ?", data.Email).
		First(&User{}).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Username already exists",
		})
		return
	}

	insert := User{
		Username: data.Username,
		Email:    data.Email,
		Password: hash.Password(data.Password),
		Phone:    data.Phone,
		Name:     data.Name,
	}

	cr.DB.Create(&insert)

	c.JSON(200, gin.H{
		"id":       insert.ID,
		"username": data.Username,
		"email":    data.Email,
		"phone":    data.Phone,
		"name":     data.Name,
	})
}
