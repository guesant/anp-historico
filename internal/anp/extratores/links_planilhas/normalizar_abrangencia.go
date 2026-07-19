package links_planilhas

import (
	"fmt"

	"github.com/guesant/anp-historico/internal/anp"
)

func normalizarAbrangencia(valor string) (anp.Abrangencia, error) {
	switch valor {
	case "brasil":
		{
			return anp.AbrangenciaBrasil, nil
		}

	case "regioes":
		{
			return anp.AbrangenciaRegioes, nil
		}

	case "estados":
		{
			return anp.AbrangenciaEstados, nil
		}

	case "municipios", "municipio":
		{
			return anp.AbrangenciaMunicipios, nil
		}

	default:
		{
			return "", fmt.Errorf("abrangência desconhecida: %q", valor)
		}
	}
}
