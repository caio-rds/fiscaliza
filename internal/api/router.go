package api

import (
	"community_voice/internal/login"
	"community_voice/internal/reports"
	user2 "community_voice/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type router struct {
}

func NewRouter() *router {
	value := router{}
	return &value
}

func (rt *router) RouteOne(db *gorm.DB) {
	r := gin.Default()
	userRouter := r.Group("/user")
	user := user2.NewRead(db)
	uCreate := user2.NewCreate(db)
	uUpdate := user2.Update(db)
	{
		userRouter.GET("/:username", user.Read)
		userRouter.POST("/", uCreate.Create)
		userRouter.PUT("/", uUpdate.UpdateUser)
	}
	loginRouter := r.Group("/login")
	tryLogin := login.NewLogin(db)
	{
		loginRouter.POST("/", tryLogin.TryLogin)
		loginRouter.GET("/protected", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": "Hello, " + username})
		})
	}

	report := r.Group("/report")
	reportCreate := reports.NewCreate(db)
	reportSearch := reports.NewRead(db)
	{
		report.POST("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			reportCreate.Create(c, username)
		})
		report.GET("/:id", reportSearch.Read)
	}

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
