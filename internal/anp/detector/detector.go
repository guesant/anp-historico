package detector

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/guesant/anp-historico/internal/anp"
	"golang.org/x/text/unicode/norm"
)

type Detector struct {
	tipoSerie   anp.TipoSerie
	abrangencia anp.Abrangencia

	indiceFisicoCabecalho int

	encontrouCabecalho bool
}

func NewDetector() *Detector {
	return &Detector{}
}

func normalizarTexto(valor string) string {
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

func normalizarLinha(linhas []string) string {
	partes := make([]string, 0, len(linhas))

	for _, coluna := range linhas {
		coluna := normalizarTexto(coluna)

		if coluna != "" {
			partes = append(partes, coluna)
		}
	}

	return strings.Join(partes, " ")
}

func ehCabecalho(linha []string) bool {
	texto := normalizarLinha(linha)

	temProduto := strings.Contains(texto, "PRODUTO")

	temPostos := strings.Contains(texto, "NUMERO DE POSTOS PESQUISADOS")

	temPeriodo := strings.Contains(texto, "MES") || strings.Contains(texto, "DATA INICIAL")

	return temProduto && temPostos && temPeriodo
}

func (d *Detector) AnalisarLinha(indiceFisicoLinha int, linha []string) {
	texto := normalizarLinha(linha)

	switch {
	case strings.Contains(texto, "INTERVALO DE TEMPO: MENSAL"):
		{
			d.tipoSerie = anp.TipoSerieMensal
		}

	case strings.Contains(texto, "INTERVALO DE TEMPO: SEMANAL"):
		{
			d.tipoSerie = anp.TipoSerieSemanal
		}
	}

	switch {
	case strings.Contains(texto, "TIPO RELATORIO: BRASIL"):
		{
			d.abrangencia = anp.AbrangenciaBrasil
		}

	case strings.Contains(texto, "TIPO RELATORIO: REGIAO"):
		{
			d.abrangencia = anp.AbrangenciaRegioes
		}

	case strings.Contains(texto, "TIPO RELATORIO: ESTADO"):
		{
			d.abrangencia = anp.AbrangenciaEstados
		}
	case strings.Contains(texto, "TIPO RELATORIO: MUNICIPIO"):
		{
			d.abrangencia = anp.AbrangenciaMunicipios
		}

	}

	if ehCabecalho(linha) {
		d.encontrouCabecalho = true
		d.indiceFisicoCabecalho = indiceFisicoLinha
	}

}

func (d *Detector) Confirmado() bool {
	return d.encontrouCabecalho && d.tipoSerie != "" && d.abrangencia != ""
}

func (d *Detector) Formato() (anp.Formato, error) {
	if !d.Confirmado() {
		return anp.Formato{}, fmt.Errorf("formato ainda não confirmado")
	}

	return anp.Formato{
		TipoSerie:             d.tipoSerie,
		Abrangencia:           d.abrangencia,
		IndiceFisicoCabecalho: d.indiceFisicoCabecalho,
	}, nil
}
