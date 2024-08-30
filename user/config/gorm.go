package config

import (
	"log"

	"github.com/susilo001/Wallet-System/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabases() {
	// setup gorm connection
	dsn := "postgres://user:password@postgres:5432/userdb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.User{})
}