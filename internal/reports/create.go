package reports

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestReport struct {
	Anonymous int    `json:"anonymous"`
	Report    string `json:"report"`
	Street    string `json:"street"`
	District  string `json:"district"`
	City      string `json:"city"`
	State     string `json:"state"`
}

type newReport struct {
	*gorm.DB
}

func NewCreate(db *gorm.DB) *newReport {
	value := newReport{
		db,
	}
	return &value
}

func (r *newReport) Create(c *gin.Context, username string) {
	var req RequestReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report := Report{
		Username:  username,
		Anonymous: req.Anonymous,
		Report:    req.Report,
		Street:    req.Street,
		District:  req.District,
		City:      req.City,
		State:     req.State,
	}

	if err := r.DB.Create(&report).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Report created"})
}
