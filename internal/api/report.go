package api

import (
	loginPkg "fiscaliza/internal/login"
	reportsPkg "fiscaliza/internal/reports"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func startReportRouter(r *gin.Engine, db *gorm.DB, login *loginPkg.Struct) {
	reports := reportsPkg.NewDb(db)
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
			reports.ReadNearest(c, username)
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
		report.GET("/websocket", login.AuthMiddleware(), func(c *gin.Context) {
			reportsPkg.Connections(c)
		})
	}
}
