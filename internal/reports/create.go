package reports

import (
	"fiscaliza/internal/models"
	"fiscaliza/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestReport struct {
	Anonymous int    `json:"anonymous"`
	Report    string `json:"report"`
	Type      string `json:"type,omitempty"`
	Street    string `json:"street"`
	District  string `json:"district"`
}

type NewReport struct {
	*gorm.DB
}

func NewCreate(db *gorm.DB) *NewReport {
	value := NewReport{
		db,
	}
	return &value
}

func (r *NewReport) Create(c *gin.Context, username string) {
	var req RequestReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coords, err := services.GetCoord(req.Street, req.District)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Type != "" {
		_, err = services.GetReportType(req.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		req.Type = "GENERIC"
	}

	report := models.Report{
		Username:  username,
		Anonymous: req.Anonymous,
		Report:    req.Report,
		Type:      req.Type,
		Street:    req.Street,
		District:  req.District,
		Lat:       coords.Latitude,
		Lon:       coords.Longitude,
	}

	if err := r.DB.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Report created", "id": report.ID})
}
