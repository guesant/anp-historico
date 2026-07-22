package cmd

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/guesant/anp-historico/internal/govbr"
	"github.com/guesant/anp-historico/internal/govbr/extratores/links_planilhas"
	"github.com/guesant/anp-historico/internal/leitor"
	"github.com/spf13/cobra"
)

var pocCmd = &cobra.Command{
	Use:   "poc",
	Short: "Proof-of-concept",

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Olá, mundo!")

		// inicio, err := anp.BuildDateISO(2005, 9, 13)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// planilha := anp.Planilha{
		// 	Tipo:        anp.TipoSerieSemanal,
		// 	Abrangencia: anp.AbrangenciaBrasil,
		// 	URL:         "https://example.org",
		// 	De:          inicio,
		// 	Ate:         nil,
		// }

		// fmt.Printf("%+v\n", planilha)
		//

		ctx := context.Background()

		planilhas, err := links_planilhas.ExtrairLinksPlanilhasDaURL(
			ctx,
			govbr.PaginaSerieHistoricaURL,
		)

		if err != nil {
			log.Fatal(err)
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		encoder.SetEscapeHTML(false)

		if err := encoder.Encode(planilhas); err != nil {
			log.Fatal(err)
		}

		var pathPlanilhas = "./planilhas"
		err = leitor.TodasAsPlanilhas(pathPlanilhas)

		if err != nil {
			log.Fatal(err)
		}

	},
}
