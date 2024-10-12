package user_address

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateRequest struct {
	ID         int     `json:"id" binding:"required"`
	Street     *string `json:"street"`
	Compliment *string `json:"compliment"`
	District   *string `json:"district"`
	City       *string `json:"city"`
	State      *string `json:"state"`
	Default    *bool   `json:"default"`
	PostalCode *string `json:"postal_code"`
	Name       *string `json:"name"`
	Lat        *string `json:"lat"`
	Lon        *string `json:"lon"`
}

func (db *Struct) Update(c *gin.Context, username string) {
	var address UpdateRequest
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAddress models.Address
	if err := db.DB.Where("id = ? AND username = ?", address.ID, username).First(&existingAddress).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	if address.Street != nil {
		existingAddress.Street = *address.Street
	}
	if address.Compliment != nil {
		existingAddress.Compliment = address.Compliment
	}
	if address.District != nil {
		existingAddress.District = *address.District
	}
	if address.City != nil {
		existingAddress.City = *address.City
	}
	if address.State != nil {
		existingAddress.State = *address.State
	}
	if address.Default != nil {
		if *address.Default {
			// Set all other addresses to non-default
			db.DB.Model(&models.Address{}).Where("username = ?", username).Update("default", false)
		}
		existingAddress.Default = *address.Default
	}
	if address.Name != nil {
		existingAddress.Name = *address.Name
	}
	if address.Lat != nil {
		existingAddress.Lat = *address.Lat
	}
	if address.Lon != nil {
		existingAddress.Lon = *address.Lon
	}

	if err := db.DB.Save(&existingAddress).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}
