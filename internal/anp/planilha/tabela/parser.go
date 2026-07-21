package tabela

import (
	"fmt"
)

type Parser struct {
	ColunasNormalizadas []ColunaTabela
}

func NewParser(colunas []string) (*Parser, error) {
	colunasNormalizadas := make([]ColunaTabela, 0, len(colunas))

	for indice, coluna := range colunas {
		colunaNormalizada, err := NormalizarColuna(coluna)

		if err != nil {
			return nil, fmt.Errorf("erro na coluna %d do cabeçalho: %w", indice+1, err)
		}

		colunasNormalizadas = append(colunasNormalizadas, colunaNormalizada)
	}

	return &Parser{
		ColunasNormalizadas: colunasNormalizadas,
	}, nil
}

func (p *Parser) ProcessarLinha(numero int, linha []string) (TabelaRegistro, error) {
	registro := TabelaRegistro{}

	for indice, colunaNormalizada := range p.ColunasNormalizadas {
		valor := ""

		if indice < len(linha) {
			valor = normalizarValor(linha[indice])
		}

		var err error

		switch colunaNormalizada {

		case ColunaTabelaMes:
			{
				registro.Mes, err = converterMes(valor)
			}
		case ColunaTabelaProduto:
			{
				registro.Produto = valor
			}
		case ColunaNumeroPostosPesquisados:
			{
				registro.PostosPesquisados, err = converterInteiro(valor)
			}
		case ColunaUnidadeDeMedida:
			{
				registro.UnidadeMedida = valor
			}
		case ColunaPrecoMedioRevenda:
			{
				registro.PrecoMedioRevenda, err = converterFloatOpcional(valor)
			}
		case ColunaDesvioPadraoRevenda:
			{
				registro.DesvioPadraoRevenda, err = converterFloatOpcional(valor)

			}
		case ColunaPrecoMinimoRevenda:
			{
				registro.PrecoMinimoRevenda, err = converterFloatOpcional(valor)

			}
		case ColunaPrecoMaximoRevenda:
			{
				registro.PrecoMaximoRevenda, err = converterFloatOpcional(valor)
			}
		case ColunaMargemMediaRevenda:
			{
				registro.MargemMediaRevenda, err = converterFloatOpcional(valor)
			}
		case ColunaCoeficienteVariacaoRevenda:
			{
				registro.CoeficienteVariacaoRevenda, err = converterFloatOpcional(valor)

			}
		case ColunaPrecoMedioDistribuicao:
			{
				registro.PrecoMedioDistribuicao, err = converterFloatOpcional(valor)

			}
		case ColunaDesvioPadraoDistribuicao:
			{
				registro.DesvioPadraoDistribuicao, err = converterFloatOpcional(valor)

			}
		case ColunaPrecoMinimoDistribuicao:
			{
				registro.PrecoMinimoDistribuicao, err = converterFloatOpcional(valor)

			}
		case ColunaPrecoMaximoDistribuicao:
			{
				registro.PrecoMaximoDistribuicao, err = converterFloatOpcional(valor)
			}
		case ColunaCoeficienteVariacaoDistribuicao:
			{
				registro.CoeficienteVariacaoDistribuicao, err = converterFloatOpcional(valor)
			}
		case ColunaRegiao:
			{
				registro.Regiao = valor
			}
		case ColunaEstado:
			{
				registro.Estado = valor
			}
		case ColunaMunicipio:
			{
				registro.Municipio = valor
			}
		case ColunaDataInicial:
			{
				registro.DataInicial, err = converterDataOpcional(valor)
			}
		case ColunaDataFinal:
			{
				registro.DataFinal, err = converterDataOpcional(valor)
			}

		default:
			{
				return TabelaRegistro{}, fmt.Errorf(
					"linha %d: coluna não reconhecida: original=%q normalizada=%q",
					numero,
					p.ColunasNormalizadas[indice],
					colunaNormalizada,
				)
			}
		}

		if err != nil {
			return TabelaRegistro{}, fmt.Errorf(
				"linha %d, coluna %q, valor %q: %w",
				numero,
				colunaNormalizada,
				valor,
				err,
			)
		}

	}

	return registro, nil
}
