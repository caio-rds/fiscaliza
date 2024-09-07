package user

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressRequest struct {
	Street     string  `json:"street"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
	Lat        string  `json:"lat"`
	Lon        string  `json:"lon"`
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

	model := models.Address{
		Username:   username,
		Street:     address.Street,
		Compliment: address.Compliment,
		District:   address.District,
		City:       "Rio de Janeiro",
		State:      "RJ",
		Lat:        address.Lat,
		Lon:        address.Lon,
	}
	db.Find(&model, "username = ?", username)
	if model.Username == "" {
		db.DB.Create(&model)
		c.JSON(200, gin.H{"message": "Address created"})
		return
	}

	if err := db.DB.Model(&models.Address{}).Where("username = ?", username).
		Updates(&model).Error; err != nil {
		c.JSON(400, gin.H{"error": "Address not found"})
		return
	}
}
