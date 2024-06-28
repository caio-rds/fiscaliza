package api

import (
	"community_voice/internal/login"
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
	read := user2.NewRead(db)
	create := user2.NewCreate(db)
	{
		userRouter.GET("/:username", read.Read)
		userRouter.POST("/", create.Create)
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

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
