package api

import (
	loginPkg "fiscaliza/internal/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startLoginRouter(r *gin.Engine, login *loginPkg.Struct) {
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("/", login.TryLogin)
		loginGroup.GET("/token", login.AuthMiddleware(), func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": username})
		})
		loginGroup.POST("/token/refresh", login.RefreshToken)
	}
}
