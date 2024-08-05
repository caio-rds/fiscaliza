package user

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
)

type SearchUser struct {
	Username string `uri:"username" binding:"required"`
}

type ResponseNoReports struct {
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	Name     string           `json:"name"`
	Address  *AddressResponse `json:"address"`
}

type AddressResponse struct {
	Street     string  `json:"street"`
	Number     string  `json:"number"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Lat        string  `json:"lat"`
	Lon        string  `json:"lon"`
}

type Response struct {
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	Name     string           `json:"name"`
	Address  *AddressResponse `json:"address"`
	Reports  []models.Report  `json:"reports"`
}

func (db *Struct) Read(c *gin.Context) {
	var search SearchUser
	if err := c.ShouldBindUri(&search); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var userAddress models.Address
	db.Find(&userAddress, "username = ?", search.Username)

	var addressResponse *AddressResponse
	if userAddress.Username != "" {
		addressResponse = &AddressResponse{
			Street:     userAddress.Street,
			Number:     userAddress.Number,
			Compliment: userAddress.Compliment,
			District:   userAddress.District,
			City:       userAddress.City,
			State:      userAddress.State,
			Lat:        userAddress.Lat,
			Lon:        userAddress.Lon,
		}
	}

	posts := c.DefaultQuery("posts", "false")
	if posts == "true" {
		var user models.User
		db.Find(&user, "username = ?", search.Username)
		var userReports []models.Report
		db.Find(&userReports, "username = ?", user.Username)
		c.JSON(200, Response{
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Name:     user.Name,
			Address:  addressResponse,
			Reports:  userReports,
		})
		return
	}

	var user models.User
	db.Find(&user, "username = ?", search.Username)
	if user.Username == "" {
		c.JSON(404, gin.H{"msg": "User not found"})
		return
	}
	c.JSON(200, ResponseNoReports{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Name:     user.Name,
		Address:  addressResponse,
	})
}
