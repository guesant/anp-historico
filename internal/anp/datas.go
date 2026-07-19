package anp

import (
	"fmt"
	"time"
)

func BuildDateISO(ano, mes, dia int) (DateISO, error) {
	data := time.Date(
		ano,
		time.Month(mes),
		dia,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	if data.Year() != ano || int(data.Month()) != mes || data.Day() != dia {
		return "", fmt.Errorf("data inválida: %04d-%02d-%02d", ano, mes, dia)
	}

	return DateISO(data.Format(time.DateOnly)), nil
}

func inicioDoAno(ano int) (DateISO, error) {
	return BuildDateISO(ano, 1, 1)
}

func fimDoAno(ano int) (DateISO, error) {
	return BuildDateISO(ano, 12, 31)
}
