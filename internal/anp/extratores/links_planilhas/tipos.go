package links_planilhas

import "github.com/guesant/anp-historico/internal/anp"

type LinkPlanilha struct {
	Tipo                anp.TipoSerie   `json:"tipo"`
	Abrangencia         anp.Abrangencia `json:"abrangencia"`
	URL                 string          `json:"url"`
	ReferenciaIntervalo string          `json:"referencia_intervalo"`
}
