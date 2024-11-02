package api

import (
	recoveryPkg "fiscaliza/internal/recovery"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func startRecoveryRouter(r *gin.Engine, db *gorm.DB) {
	rec := recoveryPkg.NewDb(db)
	recovery := r.Group("/recovery")
	{
		recovery.GET("/:username", rec.Read)
		recovery.POST("/", rec.RequestCode)
		recovery.PUT("/code", rec.ByCode)
		recovery.PUT("/similarity", rec.BySimilarity)
	}
}
