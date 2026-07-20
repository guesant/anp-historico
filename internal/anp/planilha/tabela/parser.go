package tabela

import (
	"fmt"
	"time"

	"github.com/guesant/anp-historico/internal"
)

type Parser struct {
	ColunasNormalizadas []string
}

func NewParser(colunas []string) (*Parser, error) {
	colunasNormalizadas := make([]string, 0)

	for _, coluna := range colunas {
		colunaNormalizada := internal.NormalizarTextoGenerico(coluna)
		colunasNormalizadas = append(colunasNormalizadas, colunaNormalizada)
	}

	parser := Parser{}

	parser.ColunasNormalizadas = colunasNormalizadas

	return &parser, nil
}

func (p *Parser) ProcessarLinha(numero int, linha []string) (TabelaRegistro, error) {

	registro := TabelaRegistro{}

	for indice, colunaNormalizada := range p.ColunasNormalizadas {
		valor := linha[indice]

		valor = internal.NormalizarTextoGenerico(valor)

		if valor == "-" {
			valor = ""
		}

		switch colunaNormalizada {
		case string(ColunaTabelaMes):
			{
				registro.Mes = &time.Time{}
			}

		default:
			{
				fmt.Printf("processamento não implementado: %d %s %s \n", indice, colunaNormalizada, valor)
			}
		}

	}

	return registro, nil
}
