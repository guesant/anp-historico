package tabela

import (
	"strings"

	"github.com/guesant/anp-historico/internal"
)

type ColunaTabela string

const (
	ColunaTabelaMes               ColunaTabela = "MES"
	ColunaTabelaProduto           ColunaTabela = "PRODUTO"
	ColunaNumeroPostosPesquisados ColunaTabela = "NUMERO DE POSTOS PESQUISADOS"
	UNIDADE_DE_MEDIDA             ColunaTabela = "UNIDADE DE MEDIDA"
	PRECO_MEDIO_REVENDA           ColunaTabela = "PRECO MEDIO REVENDA"
	DESVIO_PADRÃO_REVENDA         ColunaTabela = "DESVIO PADRÃO REVENDA"
	PRECO_MINIMO_REVENDA          ColunaTabela = "PRECO MINIMO REVENDA"
	PRECO_MAXIMO_REVENDA          ColunaTabela = "PRECO MAXIMO REVENDA"
	MARGEM_MEDIA_REVENDA          ColunaTabela = "MARGEM MEDIA REVENDA"
	COEF_DE_VARIACAO_REVENDA      ColunaTabela = "COEF DE VARIACAO REVENDA"
	PRECO_MEDIO_DISTRIBUICAO      ColunaTabela = "PRECO MEDIO DISTRIBUICAO"
	DESVIO_PADRÃO_DISTRIBUICAO    ColunaTabela = "DESVIO PADRÃO DISTRIBUICAO"
	PRECO_MINIMO_DISTRIBUICAO     ColunaTabela = "PRECO MINIMO DISTRIBUICAO"
	PRECO_MAXIMO_DISTRIBUICAO     ColunaTabela = "PRECO MAXIMO DISTRIBUICAO"
	COEF_DE_VARIACAO_DISTRIBUICAO ColunaTabela = "COEF DE VARIACAO DISTRIBUICAO"
	REGIAO                        ColunaTabela = "REGIAO"
	ESTADO                        ColunaTabela = "ESTADO"
	MUNICIPIO                     ColunaTabela = "MUNICÍPIO"
	ColunaDataInicial             ColunaTabela = "DATA INICIAL"
	ColunaDataFinal               ColunaTabela = "DATA FINAL"
)

func ContemColuna(colunaAlvo ColunaTabela, busca string) bool {
	return strings.Contains(busca, string(colunaAlvo))
}

func NormalizarColuna(coluna string) string {
	textoNormalizado := internal.NormalizarTextoGenerico(coluna)
	return textoNormalizado
}
