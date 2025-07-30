/*
Copyright © 2025 NAME HERE aprianfirlanda@gmail.com
*/
package cmd

import (
	"github.com/aprianfirlanda/go-server/internal/adapter/http"
	"github.com/aprianfirlanda/go-server/internal/config"
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var devMode bool

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(devMode)
		logger.InitLogger()

		port := viper.GetString("APP_PORT")
		app := http.NewFiberApp()
		http.StartServer(app, port)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.Flags().BoolVar(&devMode, "dev", false, "Run in development mode using .env")
}
