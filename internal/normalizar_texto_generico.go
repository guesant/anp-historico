package internal

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func NormalizarTextoGenerico(valor string) string {
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
