package links_planilhas

import (
	"fmt"
	"strconv"
	"time"
)

func construirIntervaloFechado(anoInicialString, anoFinalString string) (IntervaloAproximado, error) {
	anoInicial, err := strconv.Atoi(anoInicialString)

	if err != nil {
		return IntervaloAproximado{}, fmt.Errorf(
			"converter ano inicial %q: %w",
			anoInicialString,
			err,
		)
	}

	anoFinal, err := strconv.Atoi(anoFinalString)

	if err != nil {
		return IntervaloAproximado{}, fmt.Errorf(
			"converter ano final %q: %w",
			anoFinalString,
			err,
		)
	}

	if anoInicial > anoFinal {
		return IntervaloAproximado{}, fmt.Errorf(
			"intervalo invertido: %d até %d",
			anoInicial,
			anoFinal,
		)
	}

	return IntervaloAproximado{
		De:  anoInicial,
		Ate: &anoFinal,
	}, nil
}

func construirIntervaloAberto(
	anoInicialString string,
) (IntervaloAproximado, error) {
	anoInicial, err := strconv.Atoi(anoInicialString)

	if err != nil {
		return IntervaloAproximado{}, fmt.Errorf(
			"converter ano inicial %q: %w",
			anoInicialString,
			err,
		)
	}

	return IntervaloAproximado{
		De:  anoInicial,
		Ate: nil,
	}, nil
}

func normalizarIntervalo(valor string) (IntervaloAproximado, error) {
	var inicio string
	var fim *string

	if match := intervaloAnosPattern.FindStringSubmatch(valor); match != nil {
		inicio = match[1]
		fim = &match[2]
	} else if match := intervaloJaneiroPattern.FindStringSubmatch(valor); match != nil {
		inicio = match[1]
		fim = &match[2]
	} else if match := intervaloDesdeAnoPattern.FindStringSubmatch(valor); match != nil {
		inicio = match[1]
		fim = nil
	} else if match := anoUnicoPattern.FindStringSubmatch(valor); match != nil {
		inicio = match[1]

		valorFim := inicio
		fim = &valorFim
	}

	if fim != nil {
		anoFim, err := strconv.Atoi(*fim)

		if err != nil {
			return IntervaloAproximado{}, err
		}

		anoAtual := time.Now().Year()

		if anoAtual == anoFim {
			fim = nil
		}
	}

	if fim == nil {
		return construirIntervaloAberto(inicio)
	}

	return construirIntervaloFechado(inicio, *fim)
}
