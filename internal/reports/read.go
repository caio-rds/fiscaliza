package reports

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestFind struct {
	ID int `uri:"id" binding:"required"`
}

type read struct {
	*gorm.DB
}

type ReportResponse struct {
	Username  string `json:"username"`
	Anonymous int    `json:"anonymous"`
	Report    string `json:"report"`
	Street    string `json:"street"`
	District  string `json:"district"`
	City      string `json:"city"`
	State     string `json:"state"`
}

func NewRead(db *gorm.DB) *read {
	value := read{
		db,
	}
	return &value
}

func (r *read) Read(c *gin.Context) {
	var search RequestFind
	if err := c.ShouldBindUri(&search); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var report Report
	if err := r.First(&report, search.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	c.JSON(200, ReportResponse{
		Username:  report.Username,
		Anonymous: report.Anonymous,
		Report:    report.Report,
		Street:    report.Street,
		District:  report.District,
		City:      report.City,
		State:     report.State,
	})
}
