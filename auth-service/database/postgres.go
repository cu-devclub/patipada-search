package database

import (
	"fmt"
	"log"

	"auth-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	log.Println("Connecting to database....")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connected to database successfully!")

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}
