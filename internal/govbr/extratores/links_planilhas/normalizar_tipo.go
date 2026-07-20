package links_planilhas

import (
	"fmt"

	"github.com/guesant/anp-historico/internal/anp/planilha"
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
