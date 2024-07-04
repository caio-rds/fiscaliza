package models

import (
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidatePassword() (bool, error) {

	specialChars := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};:'",<>./?\\|]`)
	upperCase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)

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
