package reports

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type delete struct {
	*gorm.DB
}

func NewDelete(db *gorm.DB) *delete {
	value := delete{
		db,
	}
	return &value
}

func (d *delete) Delete(c *gin.Context, username string) {
	var search *RequestFind
	if err := c.ShouldBindUri(&search); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var report *models.Report
	if err := d.First(&report, search.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	if report.Username != username {
		c.JSON(400, gin.H{"msg": "you can only delete your own reports"})
		return
	}

	if report.DeletedAt.Valid {
		c.JSON(400, gin.H{"error": "record already deleted"})
		return
	}

	report.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}

	if err := d.Save(&report).Error; err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "record deleted"})
}
