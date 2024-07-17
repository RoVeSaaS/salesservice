package models

import (
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host           string
	Port           string
	User           string
	Password       string
	DBName         string
	SSLMode        string
	WorkOSClientId string
}

var DB *gorm.DB

func InitDB(cfg Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Customers{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
}
