package api

import (
	loginPkg "fiscaliza/internal/login"
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

	StartUserRouter(r, db)
	StartReportsRouter(r, db)

	loginRouter := r.Group("/login")
	{
		loginRouter.POST("/", login.TryLogin)
		loginRouter.GET("/token", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": username})
		})
		loginRouter.POST("/token/refresh", login.RefreshToken)
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
