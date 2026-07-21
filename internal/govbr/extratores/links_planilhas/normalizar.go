package links_planilhas

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/guesant/anp-historico/internal/anp/planilha"
	"github.com/guesant/anp-historico/internal/govbr"
)

func normalizarTipo(valor string) (planilha.TipoSerie, error) {
	switch valor {
	case "semanal":
		{
			return planilha.TipoSerieSemanal, nil
		}

	case "mensal":
		{
			return planilha.TipoSerieMensal, nil
		}

	default:
		{
			return "", fmt.Errorf("tipo de série desconhecido: %q", valor)
		}
	}
}

func normalizarAbrangencia(valor string) (planilha.Abrangencia, error) {
	switch valor {
	case "brasil":
		{
			return planilha.AbrangenciaBrasil, nil
		}

	case "regioes":
		{
			return planilha.AbrangenciaRegioes, nil
		}

	case "estados":
		{
			return planilha.AbrangenciaEstados, nil
		}

	case "municipios", "municipio":
		{
			return planilha.AbrangenciaMunicipios, nil
		}

	default:
		{
			return "", fmt.Errorf("abrangência desconhecida: %q", valor)
		}
	}
}

func normalizarIntervalo(valor string) (string, error) {
	return strings.TrimSpace(valor), nil
}

func nomeRelativo(
	rawURL string,
) (string, bool) {
	if !strings.HasPrefix(rawURL, govbr.BasePlanilhasURL) {
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
