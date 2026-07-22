package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anp-historico",
	Short: "Utilitário para consulta de histórico de preços da ANP",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(pocCmd)
}
