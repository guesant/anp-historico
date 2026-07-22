package cmd

import (
	"log"

	"github.com/guesant/anp-historico/internal/leitor"
	"github.com/spf13/cobra"
)

var pocCmd = &cobra.Command{
	Use:   "poc",
	Short: "Proof-of-concept",

	Run: func(cmd *cobra.Command, args []string) {
		var pathPlanilhas = "./planilhas"
		err := leitor.TodasAsPlanilhas(pathPlanilhas)

		if err != nil {
			log.Fatal(err)
		}
	},
}
