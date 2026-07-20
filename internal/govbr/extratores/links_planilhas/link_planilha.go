package links_planilhas

import (
	"github.com/guesant/anp-historico/internal/anp/planilha"
)

type LinkPlanilha struct {
	Tipo                planilha.TipoSerie   `json:"tipo"`
	Abrangencia         planilha.Abrangencia `json:"abrangencia"`
	URL                 string               `json:"url"`
	ReferenciaIntervalo string               `json:"referencia_intervalo"`
}
