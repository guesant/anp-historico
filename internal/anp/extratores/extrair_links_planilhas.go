package extratores

import (
	"context"
	"fmt"
	"net/url"
	"sort"

	"github.com/guesant/anp-historico/internal/anp"
	du "github.com/guesant/anp-historico/internal/anp/dom_utils"
)

func ExtrairLinksPlanilhas(
	ctx context.Context,
	paginaURL string,
) ([]anp.Planilha, error) {
	baseURL, err := url.Parse(paginaURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar URL da página %s", paginaURL)
	}

	links, err := du.ObterLinksDaURL(ctx, paginaURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", paginaURL)
	}

	planilhas := make([]anp.Planilha, 0)

	for _, href := range links {
		referencia, err := url.Parse(href)

		if err == nil {
			urlAbsoluta := baseURL.ResolveReference(referencia).String()

			planilha, reconhecida, err := anp.ClassificarPlanilha(urlAbsoluta)

			if err != nil {
				return nil, fmt.Errorf("classificar %q: %w", urlAbsoluta, err)
			}

			if reconhecida {
				planilhas = append(planilhas, planilha)
			}
		}
	}

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Tipo) < string(planilhas[j].Tipo)
	})

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Abrangencia) < string(planilhas[j].Abrangencia)
	})

	sort.Slice(planilhas, func(i, j int) bool {
		return planilhas[i].De < planilhas[j].De
	})

	return planilhas, nil
}
