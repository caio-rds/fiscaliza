package api

import (
	"fiscaliza/internal/login"
	"fiscaliza/internal/reports"
	"fiscaliza/internal/services"
	"fiscaliza/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Router struct {
}

func NewRouter() *Router {
	value := Router{}
	return &value
}

func (rt *Router) RouteOne(db *gorm.DB) {
	r := gin.Default()
	tryLogin := login.NewLogin(db)
	rep := reports.NewDb(db)
	rec := services.NewRecovery(db)
	u := user.NewDb(db)

	userRouter := r.Group("/user")
	{
		userRouter.GET("/:username", u.Read)
		userRouter.POST("/", u.Create)
		userRouter.POST("/address", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			u.UpsertAddress(c, username)
		})
		userRouter.POST("/restore/:user", u.Restore)
		userRouter.PUT("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			u.UpdateUser(c, username)
		})
		userRouter.DELETE("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			u.Delete(c, username)
		})
	}

	loginRouter := r.Group("/login")
	{
		loginRouter.POST("/", tryLogin.TryLogin)
		loginRouter.GET("/protected", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": "Hello, " + username})
		})
	}

	report := r.Group("/report")
	{
		report.POST("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			rep.Create(c, username)
		})
		report.GET("/", rep.ReadAll)
		report.GET("/nearest", rep.ReadNearest)
		report.GET("/:id", rep.Read)
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
			rep.Update(c, username, id)
		})
		report.DELETE("/:id", login.AuthMiddleware(), func(c *gin.Context) {
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
			rep.Delete(c, username)
		})
	}

	recovery := r.Group("/recovery")
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
