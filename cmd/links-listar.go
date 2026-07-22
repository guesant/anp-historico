package cmd

import (
	"context"
	"fmt"

	"github.com/guesant/anp-historico/internal/govbr"
	"github.com/guesant/anp-historico/internal/govbr/extratores/links_planilhas"
	"github.com/spf13/cobra"
)

var listarLinksCmd = &cobra.Command{
	Use:   "links-listar",
	Short: "Lista os links das planilhas da série histórica",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		planilhas, err := links_planilhas.ExtrairLinksPlanilhasDaURL(
			ctx,
			govbr.PaginaSerieHistoricaURL,
		)

		if err != nil {
			return fmt.Errorf("encontrar links: %w", err)
		}

		for _, planilha := range planilhas {
			fmt.Printf("%s\n", planilha.URL)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listarLinksCmd)
}
