package anp

import (
	"fmt"
	"strconv"
	"strings"
)

func normalizarTipo(valor string) (TipoSerie, error) {
	switch valor {
	case "semanal":
		{
			return TipoSerieSemanal, nil
		}

	case "mensal":
		{
			return TipoSerieMensal, nil
		}

	default:
		{
			return "", fmt.Errorf("tipo de série desconhecido: %q", valor)
		}
	}
}

func normalizarAbrangencia(valor string) (Abrangencia, error) {
	switch valor {
	case "brasil":
		{
			return AbrangenciaBrasil, nil
		}

	case "regioes":
		{
			return AbrangenciaRegioes, nil
		}

	case "estados":
		{
			return AbrangenciaEstados, nil
		}

	case "municipios", "municipio":
		{
			return AbrangenciaMunicipios, nil
		}

	default:
		{
			return "", fmt.Errorf("abrangência desconhecida: %q", valor)
		}
	}
}

func construirIntervaloFechado(anoInicialString, anoFinalString string) (intervalo, error) {
	anoInicial, err := strconv.Atoi(anoInicialString)

	if err != nil {
		return intervalo{}, fmt.Errorf(
			"converter ano inicial %q: %w",
			anoInicialString,
			err,
		)
	}

	anoFinal, err := strconv.Atoi(anoFinalString)

	if err != nil {
		return intervalo{}, fmt.Errorf(
			"converter ano final %q: %w",
			anoFinalString,
			err,
		)
	}

	if anoInicial > anoFinal {
		return intervalo{}, fmt.Errorf(
			"intervalo invertido: %d até %d",
			anoInicial,
			anoFinal,
		)
	}

	de, err := inicioDoAno(anoInicial)

	if err != nil {
		return intervalo{}, err
	}

	ate, err := fimDoAno(anoFinal)

	if err != nil {
		return intervalo{}, err
	}

	return intervalo{
		de:  de,
		ate: &ate,
	}, nil
}

func construirIntervaloAberto(
	anoInicialString string,
) (intervalo, error) {
	anoInicial, err := strconv.Atoi(anoInicialString)

	if err != nil {
		return intervalo{}, fmt.Errorf(
			"converter ano inicial %q: %w",
			anoInicialString,
			err,
		)
	}

	de, err := inicioDoAno(anoInicial)

	if err != nil {
		return intervalo{}, err
	}

	return intervalo{
		de:  de,
		ate: nil,
	}, nil
}

func normalizarIntervalo(valor string) (intervalo, error) {
	if match := intervaloAnosPattern.FindStringSubmatch(valor); match != nil {
		return construirIntervaloFechado(match[1], match[2])
	}

	if match := intervaloJaneiroPattern.FindStringSubmatch(valor); match != nil {
		return construirIntervaloFechado(match[1], match[2])
	}

	if match := intervaloDesdeAnoPattern.FindStringSubmatch(valor); match != nil {
		return construirIntervaloAberto(match[1])
	}

	if match := anoUnicoPattern.FindStringSubmatch(valor); match != nil {
		return construirIntervaloFechado(match[1], match[1])
	}

	return intervalo{}, fmt.Errorf(
		"intervalo desconhecido: %q",
		valor,
	)
}

func nomeRelativo(
	rawURL string,
) (string, bool) {
	if !strings.HasPrefix(rawURL, BasePlanilhasURL) {
		return "", false
	}

	nome := strings.TrimPrefix(rawURL, BasePlanilhasURL)

	nome, _, _ = strings.Cut(nome, "?")
	nome, _, _ = strings.Cut(nome, "#")

	return nome, true
}

func ClassificarPlanilha(rawURL string) (Planilha, bool, error) {
	nome, reconhecida := nomeRelativo(rawURL)
	if !reconhecida {
		return Planilha{}, false, nil
	}

	match := nomePlanilhaPattern.FindStringSubmatch(nome)

	if match == nil {
		return Planilha{}, false, nil
	}

	tipo, err := normalizarTipo(match[1])

	if err != nil {
		return Planilha{}, true, nil
	}

	abrangencia, err := normalizarAbrangencia(match[2])
	if err != nil {
		return Planilha{}, true, err
	}

	intervalo, err := normalizarIntervalo(match[3])
	if err != nil {
		return Planilha{}, true, err
	}

	return Planilha{
		Tipo:        tipo,
		Abrangencia: abrangencia,
		URL:         rawURL,
		De:          intervalo.de,
		Ate:         intervalo.ate,
	}, true, nil
}
