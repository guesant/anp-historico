package tabela

import "time"

type TabelaRegistro struct {
	Mes                        *time.Time
	Produto                    string
	PostosPesquisados          *int
	UnidadeMedida              string
	PrecoMedioRevenda          *float64
	DesvioPadraoRevenda        *float64
	PrecoMinimoRevenda         *float64
	PrecoMaximoRevenda         *float64
	MargemMediaRevenda         *float64
	CoeficienteVariacaoRevenda *float64
	PrecoMedioDistribuicao     *float64
	DesvioPadraoDistribuicao   *float64

	PrecoMinimoDistribuicao         *float64
	PrecoMaximoDistribuicao         *float64
	CoeficienteVariacaoDistribuicao *float64

	Regiao    string
	Estado    string
	Municipio string

	DataInicial *time.Time
	DataFinal   *time.Time
}
