package recovery

import (
	"fiscaliza/internal/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type CodeRequest struct {
	Username string `json:"username"`
	Type     string `json:"type"`
}

func (db *Struct) RequestCode(c *gin.Context) {

	var target string
	var req CodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFind models.User
	if err := db.Find(&userFind, "username = ?", req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userFind.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var codeExists models.Recovery
	if err := db.Find(&codeExists, "target = ? OR username = ?", target, req.Username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if codeExists.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Code already solicited"})
		return
	}
	Code := generateCode()

	var MessageId string
	if req.Type == "email" {
		target = userFind.Email
		response, err := sendEmail(&target, &Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		MessageId = *response
	} else if req.Type == "sms" {
		target = userFind.Phone
		response, err := sendSMS(&target, &Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		MessageId = *response
	}

	insert := models.Recovery{
		Target:    target,
		Username:  userFind.Username,
		Code:      Code,
		Type:      req.Type,
		MessageId: MessageId,
		ExpiresAt: time.Now().Add(time.Minute * 30),
	}

	if err := db.DB.Create(&insert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Username:  insert.Username,
		Code:      insert.Code,
		Type:      insert.Type,
		Target:    insert.Target,
		ExpiresAt: insert.ExpiresAt.Format("02/01/2006 15:04:05"),
	})

}

func generateCode() string {
	var result string
	for i := 0; i < 6; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
}
