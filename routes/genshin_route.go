package routes

import (
	"ginrest/handlers"

	"github.com/gin-gonic/gin"
)

func GenshinRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/player-search/:uid", handlers.PlayerSearchHandler)
	return r
}
