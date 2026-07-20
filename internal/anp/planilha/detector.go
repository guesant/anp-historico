package planilha

import (
	"fmt"
	"strings"

	"github.com/guesant/anp-historico/internal"
	"github.com/guesant/anp-historico/internal/anp/planilha/tabela"
)

type Detector struct {
	tipoSerie   TipoSerie
	abrangencia Abrangencia

	indiceFisicoCabecalho int

	encontrouCabecalho bool
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) AnalisarLinha(indiceFisicoLinha int, linha []string) {
	texto := internal.NormalizarLinhaGenerica(linha)

	switch {
	case strings.Contains(texto, "INTERVALO DE TEMPO: MENSAL"):
		{
			d.tipoSerie = TipoSerieMensal
		}

	case strings.Contains(texto, "INTERVALO DE TEMPO: SEMANAL"):
		{
			d.tipoSerie = TipoSerieSemanal
		}
	}

	switch {
	case strings.Contains(texto, "TIPO RELATORIO: BRASIL"):
		{
			d.abrangencia = AbrangenciaBrasil
		}

	case strings.Contains(texto, "TIPO RELATORIO: REGIAO"):
		{
			d.abrangencia = AbrangenciaRegioes
		}

	case strings.Contains(texto, "TIPO RELATORIO: ESTADO"):
		{
			d.abrangencia = AbrangenciaEstados
		}
	case strings.Contains(texto, "TIPO RELATORIO: MUNICIPIO"):
		{
			d.abrangencia = AbrangenciaMunicipios
		}

	}

	if tabela.ValidarCabecalho(linha) {
		d.encontrouCabecalho = true
		d.indiceFisicoCabecalho = indiceFisicoLinha
	}

}

func (d *Detector) Confirmado() bool {
	return d.encontrouCabecalho && d.tipoSerie != "" && d.abrangencia != ""
}

func (d *Detector) Formato() (Formato, error) {
	if !d.Confirmado() {
		return Formato{}, fmt.Errorf("formato ainda não confirmado")
	}

	return Formato{
		TipoSerie:             d.tipoSerie,
		Abrangencia:           d.abrangencia,
		IndiceFisicoCabecalho: d.indiceFisicoCabecalho,
	}, nil
}
