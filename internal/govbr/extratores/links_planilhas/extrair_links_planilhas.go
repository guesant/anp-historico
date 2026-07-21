package links_planilhas

import (
	"context"
	"fmt"
	"regexp"
	"sort"
)

func ExtrairLinksPlanilhasDaURL(
	ctx context.Context,
	paginaURL string,
) ([]LinkPlanilha, error) {
	links, err := obterLinksDaURL(ctx, paginaURL)

	if err != nil {
		return nil, fmt.Errorf("erro ao interpretar a url fornecida %q: %w", paginaURL, err)
	}

	planilhas := make([]LinkPlanilha, 0)

	for _, href := range links {
		planilha, reconhecida, err := classificarLinkPlanilha(href)

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

var (
	nomePlanilhaPattern = regexp.MustCompile(
		`^` +
			// `(?:2001-2012|semanal|mensal)/` +
			`(semanal|mensal)-` +
			`(brasil|regioes|estados|municipios?)-` +
			`(.+)\.(?:xlsx|xlsb|xls)` +
			`$`,
	)
)

func classificarLinkPlanilha(rawURL string) (LinkPlanilha, bool, error) {
	nome, reconhecida := nomeRelativo(rawURL)

	if !reconhecida {
		return LinkPlanilha{}, false, nil
	}

	match := nomePlanilhaPattern.FindStringSubmatch(nome)
	if match == nil {
		return LinkPlanilha{}, false, nil
	}

	tipo, err := normalizarTipo(match[1])
	if err != nil {
		return LinkPlanilha{}, true, nil
	}

	abrangencia, err := normalizarAbrangencia(match[2])
	if err != nil {
		return LinkPlanilha{}, true, err
	}

	referenciaIntervalo, err := normalizarIntervalo(match[3])
	if err != nil {
		return LinkPlanilha{}, true, err
	}

	return LinkPlanilha{
		Tipo:                tipo,
		Abrangencia:         abrangencia,
		URL:                 rawURL,
		ReferenciaIntervalo: referenciaIntervalo,
	}, true, nil
}
