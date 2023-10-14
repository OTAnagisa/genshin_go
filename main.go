package main

import (
	"ginrest/db"
	"ginrest/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	dbConn := db.Init()
	r := routes.Router(dbConn)
	return r
}

func main() {
	r := setupRouter()
	r.Run("localhost:8080")
}
