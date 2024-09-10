package reports

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestReport struct {
	Anonymous   int    `json:"anonymous"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description"`
	Street      string `json:"street"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func (db *StructRep) Create(c *gin.Context, username string) {
	var req RequestReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Type != "" {
		_, err := GetReportType(req.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		req.Type = "GENERIC"
	}

	report := models.Report{
		Username:    username,
		Anonymous:   req.Anonymous,
		Description: req.Description,
		Type:        req.Type,
		Street:      req.Street,
		District:    req.District,
		City:        req.City,
		State:       req.State,
		Lat:         req.Lat,
		Lon:         req.Lon,
	}

	if err := db.DB.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Report created", "id": report.ID})
}
