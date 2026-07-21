package normalizar

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func Linha(linhas []string) string {
	partes := make([]string, 0, len(linhas))

	for _, coluna := range linhas {
		coluna := Texto(coluna)

		if coluna != "" {
			partes = append(partes, coluna)
		}
	}

	return strings.Join(partes, " ")
}

func Texto(valor string) string {
	valor = strings.TrimSpace(valor)
	valor = strings.ToUpper(valor)

	decomposto := norm.NFD.String(valor)

	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Mn, r) {
			return -1
		}
		return r
	}, decomposto)
}
