package internal

import (
	"strings"
)

func NormalizarLinhaGenerica(linhas []string) string {
	partes := make([]string, 0, len(linhas))

	for _, coluna := range linhas {
		coluna := NormalizarTextoGenerico(coluna)

		if coluna != "" {
			partes = append(partes, coluna)
		}
	}

	return strings.Join(partes, " ")
}
