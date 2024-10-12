package recovery

import (
	"fiscaliza/internal/crypt"
	models "fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type byCodeRequest struct {
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type bySimilarityRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
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
