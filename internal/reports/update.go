package reports

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestUpdate struct {
	Anonymous int    `json:"anonymous"`
	Report    string `json:"report"`
	Street    string `json:"street"`
	District  string `json:"district"`
	City      string `json:"city"`
	State     string `json:"state"`
}

type update struct {
	*gorm.DB
}

func NewUpdate(db *gorm.DB) *update {
	value := update{
		db,
	}
	return &value
}

func (r *update) Update(c *gin.Context, username string, id string) {
	if username == "" {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	var req RequestUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}
	var report models.Report
	if err := r.DB.Find(&report, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Report not found",
		})
		return
	}
	if report.Username != username {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Report updated",
	})
}
