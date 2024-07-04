package services

import (
	models "community_voice/internal/models"
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

type recovery struct {
	*gorm.DB
}

func NewRecovery(db *gorm.DB) *recovery {
	value := recovery{
		db,
	}
	return &value
}

func (r *recovery) RequestCode(c *gin.Context) {
	var req CodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFind models.User
	if err := r.Find(&userFind, "email = ? OR username = ?", req.Email, req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userFind.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var codeExists models.Recovery
	if err := r.Find(&codeExists, "email = ? OR username = ?", req.Email, req.Username).Error; err != nil {
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

	if err := r.DB.Create(&insert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, RecoveryResponse{
		Email:    insert.Email,
		Username: insert.Username,
		Code:     insert.Code,
	})

}

func (r *recovery) ByCode(c *gin.Context) {
	var req *byCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var codeExists *models.Recovery
	if err := r.Find(&codeExists, "code = ?", req.Code).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if codeExists.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Code not found"})
		return
	}

	var user *models.User
	if err := r.Find(&user, "username = ?", codeExists.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = req.NewPassword

	if _, err := user.ValidatePassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = Password(user.Password)

	if err := r.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	codeExists.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}

	if err := r.Save(&codeExists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated"})

}

func (r *recovery) BySimilarity(c *gin.Context) {
	var req *bySimilarityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *models.User
	if err := r.Find(&user, "username = ?", req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := user.ValidatePassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	equalPassword := ComparePassword(user.Password, req.Password)
	if !equalPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	user.Password = Password(req.NewPassword)

	if err := r.Save(&user).Error; err != nil {
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
