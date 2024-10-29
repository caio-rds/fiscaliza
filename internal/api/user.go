package api

import (
	loginPkg "fiscaliza/internal/login"
	userPkg "fiscaliza/internal/user"
	userAdr "fiscaliza/internal/user_address"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func StartUserRouter(r *gin.Engine, db *gorm.DB) {
	user := userPkg.NewDb(db)
	address := userAdr.NewDb(db)
	login := loginPkg.NewDb(db)

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
}
