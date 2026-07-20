package links_planilhas

import (
	"fmt"

	"github.com/guesant/anp-historico/internal/anp/planilha"
)

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
