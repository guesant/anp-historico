package anp

type TipoSerie string

const (
	TipoSerieMensal  TipoSerie = "mensal"
	TipoSerieSemanal TipoSerie = "semanal"
)

type Abrangencia string

const (
	AbrangenciaBrasil     Abrangencia = "brasil"
	AbrangenciaRegioes    Abrangencia = "regioes"
	AbrangenciaEstados    Abrangencia = "estados"
	AbrangenciaMunicipios Abrangencia = "municipios"
)

type Planilha struct {
	Tipo        TipoSerie   `json:"tipo"`
	Abrangencia Abrangencia `json:"abrangencia"`
	URL         string      `json:"url"`
	De          DateISO     `json:"de"`
	Ate         *DateISO    `json:"ate"`
}

type DateISO = string

type intervalo struct {
	de  DateISO
	ate *DateISO
}
