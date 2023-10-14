package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	// DB接続設定
	dsn := "user=postgres password=postgres dbname=genshin host=localhost search_path=test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger = db.Logger.LogMode(logger.Info)

	fmt.Println("Connection Opened to Database")

	return db
}
