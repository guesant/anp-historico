package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"github.com/guesant/anp-historico/internal/anp"
)

func main() {
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

	planilhas, err := anp.ExtrairPlanilhas(
		ctx,
		anp.PaginaSerieHistoricaURL,
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
}
