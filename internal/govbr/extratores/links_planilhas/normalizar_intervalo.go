package links_planilhas

import (
	"strings"
)

func normalizarIntervalo(valor string) (string, error) {
	return strings.TrimSpace(valor), nil
}
