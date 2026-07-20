package anp

type DateISO = string

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

type Formato struct {
	TipoSerie   TipoSerie   `json:"serie"`
	Abrangencia Abrangencia `json:"brasil"`

	IndiceFisicoCabecalho int `json:"indice_fisico_cabecalho"`
}
