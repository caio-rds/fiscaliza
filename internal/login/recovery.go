package login

import (
	"fiscaliza/internal/crypt"
	models "fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type RecoveryResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Code     string `json:"code"`
}

type CodeRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type byCodeRequest struct {
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type bySimilarityRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (db *Struct) RequestCode(c *gin.Context) {
	var req CodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFind models.User
	if err := db.Find(&userFind, "email = ? OR username = ?", req.Email, req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userFind.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var codeExists models.Recovery
	if err := db.Find(&codeExists, "email = ? OR username = ?", req.Email, req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if codeExists.ID != 0 {
		c.JSON(http.StatusOK, RecoveryResponse{
			Email:    codeExists.Email,
			Username: codeExists.Username,
			Code:     codeExists.Code,
		})
		return
	}

	insert := models.Recovery{
		Email:    userFind.Email,
		Username: userFind.Username,
		Code:     generateCode(),
	}

	if err := db.DB.Create(&insert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, RecoveryResponse{
		Email:    insert.Email,
		Username: insert.Username,
		Code:     insert.Code,
	})

}

func (db *Struct) ByCode(c *gin.Context) {
	var req *byCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var codeExists *models.Recovery
	if err := db.Find(&codeExists, "code = ?", req.Code).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if codeExists.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Code not found"})
		return
	}

	var user *models.User
	if err := db.Find(&user, "username = ?", codeExists.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = req.NewPassword

	if _, err := user.ValidatePassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = crypt.Password(user.Password)

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	codeExists.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}

	if err := db.Save(&codeExists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})

}

func (db *Struct) BySimilarity(c *gin.Context) {
	var req *bySimilarityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *models.User
	if err := db.Find(&user, "username = ?", req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.ValidatePassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equalPassword := crypt.ComparePassword(user.Password, req.Password)
	if !equalPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	user.Password = crypt.Password(req.NewPassword)

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})

}

func generateCode() string {
	var result string
	for i := 0; i < 6; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
}
