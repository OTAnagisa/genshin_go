package main

import (
	"fmt"
	"ginrest/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	r := routes.GenshinRoutes()
	return r
}

func main() {
	// DB接続設定
	dsn := "user=postgres password=postgres dbname=genshin host=localhost search_path=test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	r := setupRouter()
	r.Run("localhost:8080")
}
