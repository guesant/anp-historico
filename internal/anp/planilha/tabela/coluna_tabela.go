package tabela

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/guesant/anp-historico/internal"
)

type ColunaTabela string

const (
	ColunaTabelaMes                       ColunaTabela = "MES"
	ColunaTabelaProduto                   ColunaTabela = "PRODUTO"
	ColunaNumeroPostosPesquisados         ColunaTabela = "NUMERO DE POSTOS PESQUISADOS"
	ColunaUnidadeDeMedida                 ColunaTabela = "UNIDADE DE MEDIDA"
	ColunaPrecoMedioRevenda               ColunaTabela = "PRECO MEDIO REVENDA"
	ColunaDesvioPadraoRevenda             ColunaTabela = "DESVIO PADRAO REVENDA"
	ColunaPrecoMinimoRevenda              ColunaTabela = "PRECO MINIMO REVENDA"
	ColunaPrecoMaximoRevenda              ColunaTabela = "PRECO MAXIMO REVENDA"
	ColunaMargemMediaRevenda              ColunaTabela = "MARGEM MEDIA REVENDA"
	ColunaCoeficienteVariacaoRevenda      ColunaTabela = "COEF DE VARIACAO REVENDA"
	ColunaPrecoMedioDistribuicao          ColunaTabela = "PRECO MEDIO DISTRIBUICAO"
	ColunaDesvioPadraoDistribuicao        ColunaTabela = "DESVIO PADRAO DISTRIBUICAO"
	ColunaPrecoMinimoDistribuicao         ColunaTabela = "PRECO MINIMO DISTRIBUICAO"
	ColunaPrecoMaximoDistribuicao         ColunaTabela = "PRECO MAXIMO DISTRIBUICAO"
	ColunaCoeficienteVariacaoDistribuicao ColunaTabela = "COEF DE VARIACAO DISTRIBUICAO"
	ColunaRegiao                          ColunaTabela = "REGIAO"
	ColunaEstado                          ColunaTabela = "ESTADO"
	ColunaMunicipio                       ColunaTabela = "MUNICIPIO"
	ColunaDataInicial                     ColunaTabela = "DATA INICIAL"
	ColunaDataFinal                       ColunaTabela = "DATA FINAL"
)

func ContemColuna(colunaAlvo ColunaTabela, busca string) bool {
	return strings.Contains(busca, string(colunaAlvo))
}

func NormalizarColuna(coluna string) (ColunaTabela, error) {
	textoNormaliado := internal.NormalizarTextoGenerico(coluna)

	colunaNormalizada := ColunaTabela(textoNormaliado)

	switch colunaNormalizada {

	case ColunaTabelaMes:
		{
			return colunaNormalizada, nil
		}
	case ColunaTabelaProduto:
		{
			return colunaNormalizada, nil
		}
	case ColunaNumeroPostosPesquisados:
		{
			return colunaNormalizada, nil
		}
	case ColunaUnidadeDeMedida:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMedioRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaDesvioPadraoRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMinimoRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMaximoRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaMargemMediaRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaCoeficienteVariacaoRevenda:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMedioDistribuicao:
		{
			return colunaNormalizada, nil
		}
	case ColunaDesvioPadraoDistribuicao:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMinimoDistribuicao:
		{
			return colunaNormalizada, nil
		}
	case ColunaPrecoMaximoDistribuicao:
		{
			return colunaNormalizada, nil
		}
	case ColunaCoeficienteVariacaoDistribuicao:
		{
			return colunaNormalizada, nil
		}
	case ColunaRegiao:
		{
			return colunaNormalizada, nil
		}
	case ColunaEstado:
		{
			return colunaNormalizada, nil
		}
	case ColunaMunicipio:
		{
			return colunaNormalizada, nil
		}
	case ColunaDataInicial:
		{
			return colunaNormalizada, nil
		}
	case ColunaDataFinal:
		{
			return colunaNormalizada, nil
		}

	default:
		{
			return "", fmt.Errorf("coluna não reconhecida: original=%q normalizada=%q", coluna, textoNormaliado)
		}
	}
}

func normalizarValor(valor string) string {
	valor = strings.TrimSpace(valor)

	switch strings.ToUpper(valor) {
	case "",
		"-",
		"N/D",
		"ND",
		"N/A",
		"#VALUE!",
		"#N/A",
		"#DIV/0!",
		"#REF!",
		"#NAME?",
		"#NUM!",
		"#NULL!":
		return ""

	default:
		{
			return valor
		}
	}
}

func converterMes(valor string) (*time.Time, error) {
	if valor == "" {
		return nil, nil
	}

	formatos := []string{
		"Jan-06",
		"Jan/06",
		"Jan-2006",
		"Jan/2006",
		"01/2006",
		"01-2006",
		"2006-01",
	}

	for _, formato := range formatos {
		data, err := time.Parse(formato, valor)

		if err == nil {
			return &data, nil
		}
	}

	return nil, fmt.Errorf("mês em formato desconhecido")
}

func converterInteiro(valor string) (*int, error) {
	if valor == "" {
		return nil, nil
	}

	valor = strings.TrimSpace(valor)
	valorInt, err := strconv.Atoi(valor)

	if err != nil {
		return nil, err
	}

	return &valorInt, nil
}

func converterFloatOpcional(valor string) (*float64, error) {
	if valor == "" {
		return nil, nil
	}

	valor = strings.TrimSpace(valor)

	valor = strings.ReplaceAll(valor, ",", ".")

	valorFloat, err := strconv.ParseFloat(valor, 64)

	if err != nil {
		return nil, err
	}

	return &valorFloat, nil
}

func converterDataOpcional(valor string) (*time.Time, error) {
	if valor == "" {
		return nil, nil
	}

	formatos := []string{
		"02/01/2006",
		"02/01/06",
		"2006-01-02",
		"02.01.2006",
		"02-01-06",
	}

	for _, formato := range formatos {
		data, err := time.Parse(formato, valor)

		if err == nil {
			return &data, nil
		}
	}

	return nil, fmt.Errorf("data em formato desconhecido")
}
