package address

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	Id uint `json:"id" uri:"id" binding:"required"`
}

func (db *Struct) Delete(c *gin.Context, username string) {
	var req request

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var address *models.Address
	db.Find(&address, "id = ?", req.Id)

	if address.Default {
		var addresses []*models.Address
		db.Find(&addresses, "username = ?", username)

		for _, a := range addresses {
			if a.ID == req.Id {
				continue
			}
			if err := db.DB.Model(&a).Update("default", true).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
				return
			}
		}

	}

	if err := db.DB.Delete(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
