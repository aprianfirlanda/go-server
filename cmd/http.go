/*
Copyright © 2025 NAME HERE aprianfirlanda@gmail.com
*/
package cmd

import (
	"github.com/aprianfirlanda/go-server/internal/adapter/http"
	"github.com/aprianfirlanda/go-server/internal/config"
	"github.com/aprianfirlanda/go-server/internal/database"
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var httpDevMode bool

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(httpDevMode)
		logger.InitLogger()

		database.InitPostgres()

		port := viper.GetString("APP_PORT")
		app := http.NewFiberApp()
		http.StartServer(app, port)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().BoolVar(&httpDevMode, "dev", false, "Run in development mode using .env")
}
