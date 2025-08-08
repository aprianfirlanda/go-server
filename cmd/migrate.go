package cmd

import (
	"database/sql"
	"fmt"
	"github.com/aprianfirlanda/go-server/internal/database"

	"github.com/aprianfirlanda/go-server/internal/config"
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var (
	migrationCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migration commands",
	}

	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
		Run: func(cmd *cobra.Command, args []string) {
			runMigration(func(db *sql.DB) error {
				return goose.Up(db, "db/migrations")
			})
		},
	}

	downCmd = &cobra.Command{
		Use:   "down",
		Short: "Rollback the last migration",
		Run: func(cmd *cobra.Command, args []string) {
			runMigration(func(db *sql.DB) error {
				return goose.Down(db, "db/migrations")
			})
		},
	}

	resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "Rollback all migrations to version 0",
		Run: func(cmd *cobra.Command, args []string) {
			runMigration(func(db *sql.DB) error {
				return goose.Reset(db, "db/migrations")
			})
		},
	}

	createCmd = &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new migration file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config.InitConfig(isDevMode)
			logger.InitLogger()

			if !isDevMode {
				logger.Log.Fatal("❌ You can only create migration files in development mode. Use --dev flag.")
				return
			}

			if err := goose.Create(nil, "db/migrations", args[0], "sql"); err != nil {
				logger.Log.WithError(err).Fatal("Failed to create migration")
			}

			fmt.Println("✅ Migration file created in db/migrations")
		},
	}
)

func init() {
	rootCmd.AddCommand(migrationCmd)

	migrationCmd.AddCommand(upCmd)
	migrationCmd.AddCommand(downCmd)
	migrationCmd.AddCommand(resetCmd)
	migrationCmd.AddCommand(createCmd)
}

func runMigration(action func(*sql.DB) error) {
	config.InitConfig(isDevMode)
	logger.InitLogger()

	database.InitPostgres()
	db, _ := database.DB.DB()

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Log.WithError(err).Fatal("❌ Failed to set dialect")
	}

	if err := action(db); err != nil {
		logger.Log.WithError(err).Fatal("❌ Migration command failed")
	}
}
