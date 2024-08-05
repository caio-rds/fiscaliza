package user

import (
	"fiscaliza/internal/models"
	"fiscaliza/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
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
	phone := regexp.MustCompile(`^\(\d{2}\) \d{5}-\d{4}$`)

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false, fmt.Errorf("invalid email")
	}
	if !phone.MatchString(u.Phone) {
		return false, fmt.Errorf("phone number must have (xx)xxxxx-xxxx format")
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

func (db *Struct) Create(c *gin.Context) {
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

	if err := db.DB.Where("username = ?", data.Username).Or("email = ?", data.Email).
		First(&models.User{}).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "Username already exists",
		})
		return
	}

	insert := models.User{
		Username: data.Username,
		Email:    data.Email,
		Password: services.Password(data.Password),
		Phone:    data.Phone,
		Name:     data.Name,
	}

	if err := db.DB.Create(&insert).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id":       insert.ID,
		"username": data.Username,
		"email":    data.Email,
		"phone":    data.Phone,
		"name":     data.Name,
	})
}
