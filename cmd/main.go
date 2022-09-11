package main

import (
	"log"
	"os"

	"github.com/raytr/sample-project-p/config"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config

	rootCmd = &cobra.Command{
		SilenceErrors: true,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serverCmd)
}

func initConfig() {
	cfg = config.Load()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		log.Printf("Succeeded")
	}
}
