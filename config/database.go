package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	port := config.GetInt("DB_PORT")
	database := config.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error getting sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	return db
}
