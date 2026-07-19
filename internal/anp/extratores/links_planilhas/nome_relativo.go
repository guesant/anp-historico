package links_planilhas

import (
	"net/url"
	"path"
	"strings"

	"github.com/guesant/anp-historico/internal/anp"
)

func nomeRelativo(
	rawURL string,
) (string, bool) {
	if !strings.HasPrefix(rawURL, anp.BasePlanilhasURL) {
		return "", false
	}

	return extrairUltimoPath(rawURL)
}

func extrairUltimoPath(rawURL string) (string, bool) {
	parsed, err := url.Parse(rawURL)

	if err != nil {
		return "", false
	}

	if parsed.Path == "" || strings.HasSuffix(parsed.Path, "/") {
		return "", false
	}

	nome := path.Base(parsed.Path)

	if nome == "." || nome == "/" || nome == "" {
		return "", false
	}

	return nome, true
}
