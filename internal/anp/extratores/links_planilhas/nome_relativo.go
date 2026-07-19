package links_planilhas

import (
	"strings"

	"github.com/guesant/anp-historico/internal/anp"
)

func nomeRelativo(
	rawURL string,
) (string, bool) {
	if !strings.HasPrefix(rawURL, anp.BasePlanilhasURL) {
		return "", false
	}

	nome := strings.TrimPrefix(rawURL, anp.BasePlanilhasURL)

	nome, _, _ = strings.Cut(nome, "?")
	nome, _, _ = strings.Cut(nome, "#")

	return nome, true
}
