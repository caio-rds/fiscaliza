package address

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
)

type addressResponse struct {
	Id         uint    `json:"id"`
	Street     string  `json:"street"`
	Compliment *string `json:"compliment"`
	District   string  `json:"district"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postal_code"`
	Name       string  `json:"name"`
	Default    bool    `json:"default"`
	Lat        string  `json:"lat"`
	Lon        string  `json:"lon"`
}

type params struct {
	Id *uint `uri:"id"`
}

func (db *Struct) Read(c *gin.Context, username *string) {

	var p params
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var address []*models.Address
	db.Find(&address, "username = ?", username)

	var userAddresses []*addressResponse
	var responseDefault *addressResponse

	for _, adr := range address {
		response := addressResponse{
			Id:         adr.ID,
			Street:     adr.Street,
			Compliment: adr.Compliment,
			District:   adr.District,
			City:       adr.City,
			State:      adr.State,
			Name:       adr.Name,
			Default:    adr.Default,
			Lat:        adr.Lat,
			Lon:        adr.Lon,
		}

		if p.Id != nil && *p.Id == adr.ID {
			responseDefault = &response
		} else {
			userAddresses = append(userAddresses, &response)
		}
	}

	if responseDefault != nil {
		c.JSON(200, &responseDefault)
		return
	}

	c.JSON(200, &userAddresses)

}
