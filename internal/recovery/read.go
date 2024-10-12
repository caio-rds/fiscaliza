package recovery

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Username  string `json:"username"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	Target    string `json:"target,omitempty"`
	ExpiresAt string `json:"expires_at"`
}

func (db *Struct) Read(c *gin.Context) {

	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var recovery models.Recovery
	if err := db.Find(&recovery, "username = ?", username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if recovery.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Recovery solicitation found"})
		return
	}

	c.JSON(http.StatusOK, Response{
		Username:  recovery.Username,
		Code:      recovery.Code,
		Type:      recovery.Type,
		ExpiresAt: recovery.ExpiresAt.Format("02/01/2006 15:04:05"),
	})
}
