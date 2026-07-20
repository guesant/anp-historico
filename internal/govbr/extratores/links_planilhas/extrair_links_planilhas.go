package links_planilhas

import (
	"context"
	"fmt"
	"sort"

	du "github.com/guesant/anp-historico/internal/html_utils"
)

func ExtrairLinksPlanilhasDaURL(
	ctx context.Context,
	paginaURL string,
) ([]LinkPlanilha, error) {
	links, err := du.ObterLinksDaURL(ctx, paginaURL)

	if err != nil {
		return nil, fmt.Errorf("erro ao interpretar a url fornecida %q: %w", paginaURL, err)
	}

	planilhas := make([]LinkPlanilha, 0)

	for _, href := range links {
		planilha, reconhecida, err := ClassificarLinkPlanilha(href)

		if err != nil {
			return nil, fmt.Errorf("erro ao tentar classificar uma planilha presente na página %q: %w", href, err)
		}

		if reconhecida {
			planilhas = append(planilhas, planilha)
		}
	}

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Tipo) < string(planilhas[j].Tipo)
	})

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Abrangencia) < string(planilhas[j].Abrangencia)
	})

	return planilhas, nil
}
