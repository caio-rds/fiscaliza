package api

import (
	"community_voice/internal/login"
	"community_voice/internal/reports"
	"community_voice/internal/services"
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
	uRead := user2.NewRead(db)
	uCreate := user2.NewCreate(db)
	uRestore := user2.NewRestore(db)
	uUpdate := user2.Update(db)
	uDelete := user2.NewDelete(db)
	{
		userRouter.GET("/:username", uRead.Read)
		userRouter.POST("/", uCreate.Create)
		userRouter.POST("/restore/:user", uRestore.Restore)
		userRouter.PUT("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			uUpdate.UpdateUser(c, username)
		})
		userRouter.DELETE("/:username", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			uDelete.Delete(c, username)
		})
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
	reportUpdate := reports.NewUpdate(db)
	{
		report.POST("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			reportCreate.Create(c, username)
		})
		report.GET("/", reportSearch.ReadAll)
		report.GET("/nearest", reportSearch.ReadNearest)
		report.GET("/:id", reportSearch.Read)
		report.PUT("/:id", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			reportUpdate.Update(c, username, id)
		})
	}

	recovery := r.Group("/recovery")
	rec := services.NewRecovery(db)
	{
		recovery.POST("/", rec.RequestCode)
		recovery.POST("/code", rec.ByCode)
		recovery.POST("/similarity", rec.BySimilarity)
	}

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
