package reports

import (
	"fiscaliza/internal/models"
	"fiscaliza/internal/services"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type RequestFind struct {
	ID int `uri:"id" binding:"required"`
}

type ReportResponse struct {
	Id          uint      `json:"id"`
	Username    string    `json:"username"`
	Anonymous   int       `json:"anonymous"`
	Type        *string   `json:"type"`
	Description string    `json:"description"`
	Street      string    `json:"street"`
	District    string    `json:"district"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	Solved      int       `json:"solved"`
	Lat         string    `json:"lat"`
	Lon         string    `json:"lon"`
}

type Filters struct {
	Street   string `json:"street"`
	District string `json:"district"`
	Reverse  bool   `json:"created" default:"false"`
}

func (db *StructRep) Read(c *gin.Context) {
	var search *RequestFind
	if err := c.ShouldBindUri(&search); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var report *models.Report
	if err := db.First(&report, search.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	c.JSON(200, &ReportResponse{
		Id:          report.ID,
		Username:    report.Username,
		Anonymous:   report.Anonymous,
		Description: report.Description,
		Type:        services.GetReportTypeName(report.Type),
		Street:      report.Street,
		District:    report.District,
		City:        report.City,
		State:       report.State,
		CreatedAt:   report.CreatedAt,
		Lat:         report.Lat,
		Lon:         report.Lon,
	})
}

func (db *StructRep) ReadAll(c *gin.Context) {
	var filters Filters
	var url = c.Request.URL.Query()

	if len(url) > 0 {
		if url["street"] != nil {
			filters.Street = url["street"][0]
		}
		if url["district"] != nil {
			filters.District = url["district"][0]
		}
		if url["reverse"] != nil {
			filters.Reverse = true
		}
	}

	var reports *[]models.Report
	if filters.Street == "" && filters.District == "" {
		if err := db.Find(&reports).Error; err != nil {
			c.JSON(404, gin.H{"error": "record not found"})
			return
		}
	} else {
		if filters.Street != "" && filters.District != "" {
			if err := db.Where("street = ? AND district = ?", filters.Street, filters.District).Find(&reports).Error; err != nil {
				c.JSON(404, gin.H{"error": "There is no report with this street and district"})
				return
			}
		} else if filters.Street != "" {
			if err := db.Where("street = ?", filters.Street).Find(&reports).Error; err != nil {
				c.JSON(404, gin.H{"error": "record not found"})
				return
			}
		} else if filters.District != "" {
			if err := db.Where("district = ?", filters.District).Find(&reports).Error; err != nil {
				c.JSON(404, gin.H{"error": "record not found"})
				return
			}
		}
	}

	if len(*reports) == 0 {
		c.JSON(404, gin.H{"error": "no reports found"})
		return
	}

	var response []ReportResponse
	for _, report := range *reports {
		if report.Anonymous == 1 {
			response = append(response, ReportResponse{
				Id:          report.ID,
				Username:    "not available",
				Anonymous:   report.Anonymous,
				Description: report.Description,
				Type:        services.GetReportTypeName(report.Type),
				Street:      report.Street,
				District:    report.District,
				City:        report.City,
				State:       report.State,
				CreatedAt:   report.CreatedAt,
				Lat:         report.Lat,
				Lon:         report.Lon,
			})
			continue
		}
		response = append(response, ReportResponse{
			Id:          report.ID,
			Username:    report.Username,
			Anonymous:   report.Anonymous,
			Description: report.Description,
			Type:        services.GetReportTypeName(report.Type),
			Street:      report.Street,
			District:    report.District,
			City:        report.City,
			State:       report.State,
			CreatedAt:   report.CreatedAt,
			Lat:         report.Lat,
			Lon:         report.Lon,
		})
	}
	if !filters.Reverse {
		for i, j := 0, len(response)-1; i < j; i, j = i+1, j-1 {
			response[i], response[j] = response[j], response[i]
		}
	}

	c.JSON(200, response)
}

type NearestReports struct {
	Lat   string  `json:"lat"`
	Lon   string  `json:"lon"`
	Range float64 `json:"range" default:"1.0"`
}

func (db *StructRep) ReadNearest(c *gin.Context) {
	var currentCoords NearestReports
	var url = c.Request.URL.Query()

	if len(url) > 0 {
		if url["lat"] != nil {
			currentCoords.Lat = url["lat"][0]
		}
		if url["lon"] != nil {
			currentCoords.Lon = url["lon"][0]
		}
		if url["range"] != nil {
			currentCoords.Range, _ = strconv.ParseFloat(url["range"][0], 64)
		}
	}

	var reports *[]models.Report
	if err := db.Find(&reports).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	var response []ReportResponse
	for _, report := range *reports {
		distance, err := services.GetDistance(currentCoords.Lat, currentCoords.Lon, report.Lat, report.Lon)
		if err != nil {
			c.JSON(404, gin.H{"error": "no coordinates found"})
			return
		}
		if distance.Distance <= currentCoords.Range {
			if report.Anonymous == 1 {
				response = append(response, ReportResponse{
					Id:          report.ID,
					Username:    "not available",
					Anonymous:   report.Anonymous,
					Description: report.Description,
					Type:        services.GetReportTypeName(report.Type),
					Street:      report.Street,
					District:    report.District,
					City:        report.City,
					State:       report.State,
					CreatedAt:   report.CreatedAt,
					Lat:         report.Lat,
					Lon:         report.Lon,
				})
				continue
			}
			response = append(response, ReportResponse{
				Id:          report.ID,
				Username:    report.Username,
				Anonymous:   report.Anonymous,
				Description: report.Description,
				Type:        services.GetReportTypeName(report.Type),
				Street:      report.Street,
				District:    report.District,
				City:        report.City,
				State:       report.State,
				CreatedAt:   report.CreatedAt,
				Lat:         report.Lat,
				Lon:         report.Lon,
			})
		}
	}

	if len(response) == 0 {
		c.JSON(404, gin.H{"error": "no reports found in your area"})
		return
	}

	for i, j := 0, len(response)-1; i < j; i, j = i+1, j-1 {
		response[i], response[j] = response[j], response[i]
	}
	c.JSON(200, response)

}
