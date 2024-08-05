package user

import (
	"fiscaliza/internal/models"
	"fiscaliza/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AddressRequest struct {
	Street     string  `json:"street"`
	Number     string  `json:"number"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
}

func (db *Struct) UpsertAddress(c *gin.Context, username string) {
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}
	var address AddressRequest
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	coords, err := services.GetCoord(fmt.Sprintf("%s, %s", address.Street, address.Number), address.District)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var model models.Address
	db.Find(&model, "username = ?", username)
	if model.Username == "" {
		model = models.Address{
			Username:   username,
			Street:     address.Street,
			Number:     address.Number,
			Compliment: address.Compliment,
			District:   strings.ToUpper(string(address.District[0])) + strings.ToLower(address.District[1:]),
			City:       "Rio de Janeiro",
			State:      "RJ",
			Lat:        coords.Latitude,
			Lon:        coords.Longitude,
		}
		db.DB.Create(&model)
		c.JSON(200, gin.H{"message": "Address created"})
		return
	}

	if err := db.DB.Model(&models.Address{}).Where("username = ?", username).
		Updates(models.Address{
			Street:     address.Street,
			Number:     address.Number,
			Compliment: address.Compliment,
			District:   strings.ToUpper(string(address.District[0])) + strings.ToLower(address.District[1:]),
			City:       "Rio de Janeiro",
			State:      "RJ",
			Lat:        coords.Latitude,
			Lon:        coords.Longitude,
		}).Error; err != nil {
		c.JSON(400, gin.H{"error": "Address not found"})
		return
	}
}
