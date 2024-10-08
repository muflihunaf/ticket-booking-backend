package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/muflihunaf/ticket-booking-v1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	url := fmt.Sprintf(`
		host=%s user=%s dbname=%s password=%s sslmode=%s port=5432
	`, config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Unable to connect to database %e", err)
	}
	log.Info("connected to database")

	if err := DBMigrator(db); err != nil {
		log.Fatal("Unable to migrate the table %e", err)
	}

	return db
}
