package reports

import (
	"fiscaliza/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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
	Distance    *string   `json:"distance"`
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

	if report.Anonymous == 1 {
		report.Username = "not available"
	}

	c.JSON(200, &ReportResponse{
		Id:          report.ID,
		Username:    report.Username,
		Anonymous:   report.Anonymous,
		Description: report.Description,
		Type:        GetReportTypeName(report.Type),
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
		newResponse := ReportResponse{
			Id:          report.ID,
			Username:    report.Username,
			Anonymous:   report.Anonymous,
			Description: report.Description,
			Type:        GetReportTypeName(report.Type),
			Street:      report.Street,
			District:    report.District,
			City:        report.City,
			State:       report.State,
			CreatedAt:   report.CreatedAt,
			Lat:         report.Lat,
			Lon:         report.Lon,
		}
		if newResponse.Anonymous == 1 {
			newResponse.Username = "not available"
		}

		response = append(response, newResponse)
	}
	if !filters.Reverse {
		for i, j := 0, len(response)-1; i < j; i, j = i+1, j-1 {
			response[i], response[j] = response[j], response[i]
		}
	}

	c.JSON(200, response)
}

type NearestReports struct {
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Range       float64 `json:"range" default:"0"`
	HomeAddress bool    `json:"home_address" default:"false"`
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
		distance, err := getDistance(currentCoords.Lat, currentCoords.Lon, report.Lat, report.Lon)
		if err != nil {
			c.JSON(404, gin.H{"error": "no coordinates found"})
			return
		}
		var reportDistance = fmt.Sprintf("%.3f", distance.Distance)

		newResponse := ReportResponse{
			Id:          report.ID,
			Username:    report.Username,
			Anonymous:   report.Anonymous,
			Description: report.Description,
			Type:        GetReportTypeName(report.Type),
			Distance:    &reportDistance,
			Street:      report.Street,
			District:    report.District,
			City:        report.City,
			State:       report.State,
			CreatedAt:   report.CreatedAt,
			Lat:         report.Lat,
			Lon:         report.Lon,
		}
		if newResponse.Anonymous == 1 {
			newResponse.Username = "not available"
		}

		if url["range"] != nil {
			if distance.Distance <= currentCoords.Range {
				response = append(response, newResponse)
			}
		} else {
			response = append(response, newResponse)
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

type Distance struct {
	Distance   float64 `json:"distance"`
	Latitude1  string  `json:"latitude1"`
	Longitude1 string  `json:"longitude1"`
	Latitude2  string  `json:"latitude2"`
	Longitude2 string  `json:"longitude2"`
}

func toRadians(degrees string) (float64, error) {
	degreesFloat, err := strconv.ParseFloat(degrees, 64)
	if err != nil {
		return 0, err
	}
	return degreesFloat * (math.Pi / 180), nil
}

func getDistance(lat1 string, lon1 string, lat2 string, lon2 string) (*Distance, error) {
	const EarthRadius = 6371

	lat1Rad, err := toRadians(lat1)
	if err != nil {
		return nil, err
	}
	lon1Rad, err := toRadians(lon1)
	if err != nil {
		return nil, err
	}
	lat2Rad, err := toRadians(lat2)
	if err != nil {
		return nil, err
	}
	lon2Rad, err := toRadians(lon2)
	if err != nil {
		return nil, err
	}

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadius * c

	return &Distance{
		Distance:   distance,
		Latitude1:  lat1,
		Longitude1: lon1,
		Latitude2:  lat2,
		Longitude2: lon2,
	}, nil
}

func (db *StructRep) ReportsByUser(c *gin.Context, username string) {

	var reports *[]models.Report
	if err := db.Where("username = ?", username).Find(&reports).Error; err != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}

	if len(*reports) == 0 {
		c.JSON(404, gin.H{"error": "no reports found"})
		return
	}

	var response []ReportResponse
	for _, report := range *reports {
		newResponse := ReportResponse{
			Id:          report.ID,
			Username:    report.Username,
			Anonymous:   report.Anonymous,
			Description: report.Description,
			Type:        GetReportTypeName(report.Type),
			Street:      report.Street,
			District:    report.District,
			City:        report.City,
			State:       report.State,
			CreatedAt:   report.CreatedAt,
			Lat:         report.Lat,
			Lon:         report.Lon,
		}

		response = append(response, newResponse)
	}

	c.JSON(200, response)
}
