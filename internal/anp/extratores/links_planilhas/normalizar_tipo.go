package links_planilhas

import (
	"fmt"

	"github.com/guesant/anp-historico/internal/anp"
)

func normalizarTipo(valor string) (anp.TipoSerie, error) {
	switch valor {
	case "semanal":
		{
			return anp.TipoSerieSemanal, nil
		}

	case "mensal":
		{
			return anp.TipoSerieMensal, nil
		}

	default:
		{
			return "", fmt.Errorf("tipo de série desconhecido: %q", valor)
		}
	}
}
