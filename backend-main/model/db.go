package model

import (
	"backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	switch config.AppConfig.DBType {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(config.AppConfig.SQLitePath))
	case "postgres":
		DB, err = gorm.Open(postgres.Open(config.AppConfig.PostgresDSN))
	default:
		panic("Unsupported db backend! Valid options: sqlite, postgres")
	}
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Nexus{}, &User{})
	if err != nil {
		panic(err)
	}
}
