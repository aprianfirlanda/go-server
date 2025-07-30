package config

import (
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/spf13/viper"
)

// InitConfig initializes configuration using Viper.
// If devMode is true, it loads from .env; otherwise, from env vars.
func InitConfig(devMode bool) {
	log := logger.TempLogger()

	viper.AutomaticEnv()

	if devMode {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatal("Error reading .env file")
		}
		log.Info("Loaded config from .env file")
	} else {
		log.Info("Loaded config from environment")
	}
}
