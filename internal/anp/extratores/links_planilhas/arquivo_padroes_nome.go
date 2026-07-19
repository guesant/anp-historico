package links_planilhas

import "regexp"

var (
	nomePlanilhaPattern = regexp.MustCompile(
		`^(?:2001-2012|semanal|mensal)/` +
			`(semanal|mensal)-` +
			`(brasil|regioes|estados|municipios?)-` +
			`(.+)\.(?:xlsx|xlsb|xls)$`,
	)

	intervaloAnosPattern = regexp.MustCompile(
		`^(\d{4})(?:-a-|-|_a_)(\d{4})$`,
	)

	intervaloDesdeAnoPattern = regexp.MustCompile(`^desde-(?:jan)?(\d{4})$`)

	intervaloJaneiroPattern = regexp.MustCompile(
		`^jan(\d{4})-(\d{4})$`,
	)

	anoUnicoPattern = regexp.MustCompile(
		`^(\d{4})$`,
	)
)
