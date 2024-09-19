package api

import (
	loginPkg "fiscaliza/internal/login"
	reportsPkg "fiscaliza/internal/reports"
	userPkg "fiscaliza/internal/user"
	userAdr "fiscaliza/internal/user_address"
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

func (rt *Router) StartRouter(db *gorm.DB) {
	r := gin.Default()
	login := loginPkg.NewDb(db)
	reports := reportsPkg.NewDb(db)
	user := userPkg.NewDb(db)
	address := userAdr.NewDb(db)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "allow-control-allow-origin",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRouter := r.Group("/user")
	{
		userRouter.GET("/:username", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.Param("username")
			if username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			user.Read(c, "")
		})
		userRouter.GET("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			user.Read(c, username)
		})
		userRouter.POST("/", user.Create)
		userRouter.POST("/restore/:user", user.Restore)
		userRouter.PUT("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			user.UpdateUser(c, username)
		})
		userRouter.DELETE("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			user.Delete(c, username)
		})
		userRouter.GET("/address", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			address.Read(c, &username)
		})
		userRouter.POST("/address", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			address.Create(c, username)
		})
		userRouter.PUT("/address", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			address.Update(c, username)
		})
		userRouter.DELETE("/address/:id", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			address.Delete(c, username)
		})
	}

	loginRouter := r.Group("/login")
	{
		loginRouter.POST("/", login.TryLogin)
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
			reports.Create(c, username)
		})
		report.GET("/all", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			reports.ReadAll(c)
		})
		report.GET("/all/own", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.Param("username")
			if username == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			reports.ReportsByUser(c, username)
		})
		report.GET("/", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			reports.ReadNearest(c)
		})
		report.GET("/:id", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			reports.Read(c)
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
			reports.Update(c, username, id)
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
			reports.Delete(c, username)
		})
		report.GET("/types", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			if username == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.JSON(http.StatusOK, reportsPkg.GetReportTypes())
		})

	}

	recovery := r.Group("/recovery")
	{
		recovery.POST("/", login.RequestCode)
		recovery.POST("/code", login.ByCode)
		recovery.POST("/similarity", login.BySimilarity)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy"})
	})

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
