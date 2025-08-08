/*
Copyright © 2025 NAME HERE aprianfirlanda@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-server",
		Short: "Backend application that create with GO",
	}
	isDevMode bool
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	checkDevMode()
	fmt.Println(isDevMode)
}

func checkDevMode() {
	path, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: cannot get executable path: %v\n", err)
		return
	}
	isDevMode = strings.Contains(path, "go-build")
}
