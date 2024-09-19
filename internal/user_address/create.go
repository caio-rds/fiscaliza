package user_address

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressRequest struct {
	Street     string  `json:"street"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Default    bool    `json:"default"`
	PostalCode string  `json:"postal_code"`
	Name       string  `json:"name"`
	Lat        string  `json:"lat"`
	Lon        string  `json:"lon"`
}

func (db *Struct) Create(c *gin.Context, username string) {
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

	var addresses *[]models.Address
	db.Find(&addresses, "username = ?", username)

	if len(*addresses) == 0 {
		address.Default = true
	}

	if len(*addresses) >= 4 {
		c.JSON(400, gin.H{"error": "You can't have more than 4 addresses"})
		return
	}

	for _, adr := range *addresses {

		if address.Name == adr.Name {
			c.JSON(400, gin.H{"error": "Address name already exists"})
			return
		}

		if address.Default {
			adr.Default = false
			db.DB.Save(&adr)
		}

	}

	insert := models.Address{
		Username:   username,
		Street:     address.Street,
		Compliment: address.Compliment,
		District:   address.District,
		City:       address.City,
		State:      address.State,
		Default:    address.Default,
		PostalCode: address.PostalCode,
		Name:       address.Name,
		Lat:        address.Lat,
		Lon:        address.Lon,
	}

	db.DB.Create(&insert)
	c.JSON(200, gin.H{"message": "Address created"})
}
