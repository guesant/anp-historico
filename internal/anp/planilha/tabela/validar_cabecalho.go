package tabela

import (
	"github.com/guesant/anp-historico/internal"
)

func ValidarCabecalho(linha []string) bool {
	texto := internal.NormalizarLinhaGenerica(linha)

	temProduto := ContemColuna(ColunaTabelaProduto, texto)
	temPostos := ContemColuna(ColunaNumeroPostosPesquisados, texto)
	temPeriodo := ContemColuna(ColunaTabelaMes, texto) || ContemColuna(ColunaDataInicial, texto)

	cabecalho := temProduto && temPostos && temPeriodo

	return cabecalho
}
