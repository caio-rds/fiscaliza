package user

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
)

type SearchUser struct {
	Username string `uri:"username"`
}

type AddressResponse struct {
	Id         uint    `json:"id"`
	Street     string  `json:"street"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postal_code"`
	Default    bool    `json:"default"`
	Name       string  `json:"name"`
	Lat        string  `json:"lat"`
	Lon        string  `json:"lon"`
}

type Response struct {
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Phone    string             `json:"phone"`
	Name     string             `json:"name"`
	Address  []*AddressResponse `json:"address"`
}

func (db *Struct) Read(c *gin.Context, username string) {
	var search SearchUser

	if username == "" {
		if err := c.ShouldBindUri(&search); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
	} else {
		search.Username = username
	}

	var userAddress []*models.Address
	db.Find(&userAddress, "username = ?", search.Username)

	var addressResponse []*AddressResponse
	adrDefault := c.DefaultQuery("address_default", "false")

	for _, address := range userAddress {
		if adrDefault == "true" && !address.Default {
			continue
		}
		addressResponse = append(addressResponse, &AddressResponse{
			Id:         address.ID,
			Street:     address.Street,
			Compliment: address.Compliment,
			District:   address.District,
			City:       address.City,
			State:      address.State,
			PostalCode: address.PostalCode,
			Default:    address.Default,
			Name:       address.Name,
			Lat:        address.Lat,
			Lon:        address.Lon,
		})
	}

	var user models.User
	db.Find(&user, "username = ?", search.Username)
	if user.Username == "" {
		c.JSON(404, gin.H{"msg": "User not found"})
		return
	}

	c.JSON(200, Response{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Name:     user.Name,
		Address:  addressResponse,
	})
}
