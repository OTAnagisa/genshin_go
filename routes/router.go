package routes

import (
	"ginrest/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(dbConn *gorm.DB) *gin.Engine {
	r := gin.Default()

	languageHandler := controllers.LanguageHandler{Db: dbConn}
	r.GET("/language", languageHandler.GetAll)

	r.GET("/player-search/:uid", controllers.PlayerSearchHandler)
	return r
}
