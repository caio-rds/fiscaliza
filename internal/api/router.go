package api

import (
	"fiscaliza/internal/login"
	"fiscaliza/internal/reports"
	"fiscaliza/internal/services"
	"fiscaliza/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "ngrok-skip-browser-warning", "allow-control-allow-origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		loginRouter.GET("/token", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": username})
		})
		loginRouter.POST("/token/refresh", login.RefreshToken)
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
		report.GET("/all", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			rep.ReadAll(c)
		})
		report.GET("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			rep.ReadNearest(c)
		})
		report.GET("/:id", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			rep.Read(c)
		})
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
		report.GET("/types", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.JSON(http.StatusOK, services.GetReportTypes())
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
