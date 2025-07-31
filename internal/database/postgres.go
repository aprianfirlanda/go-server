package database

import (
	"fmt"
	"time"

	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger" // 👈 alias this to avoid name clash
)

var DB *gorm.DB

func InitPostgres() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_SSLMODE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info), // ✅ no clash
	})
	if err != nil {
		logger.Log.WithError(err).Fatal("❌ Failed to connect to Postgres")
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Log.Info("✅ Connected to Postgres")
}
