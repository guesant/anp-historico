package links_planilhas

import "github.com/guesant/anp-historico/internal/anp"

type LinkPlanilha struct {
	Tipo                anp.TipoSerie       `json:"tipo"`
	Abrangencia         anp.Abrangencia     `json:"abrangencia"`
	URL                 string              `json:"url"`
	IntervaloAproximado IntervaloAproximado `json:"intervaloAproximado"`
}

type IntervaloAproximado struct {
	De  int  `json:"de"`
	Ate *int `json:"ate"`
}
