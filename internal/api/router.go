package api

import (
	loginPkg "fiscaliza/internal/login"
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

func (rt *Router) StartRouter(db *gorm.DB) {
	r := gin.Default()
	login := loginPkg.NewDb(db)

	//r.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders: []string{
	//		"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin",
	//		"Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "allow-control-allow-origin",
	//	},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	configureCORS(r)
	startUserRouter(r, db, login)
	startReportRouter(r, db, login)
	startLoginRouter(r, login)
	startRecoveryRouter(r, db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy"})
	})

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("panic: %v", err)
		return
	}
}
