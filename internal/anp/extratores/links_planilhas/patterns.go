package links_planilhas

import "regexp"

var (
	nomePlanilhaPattern = regexp.MustCompile(
		`^` +
			// `(?:2001-2012|semanal|mensal)/` +
			`(semanal|mensal)-` +
			`(brasil|regioes|estados|municipios?)-` +
			`(.+)\.(?:xlsx|xlsb|xls)` +
			`$`,
	)
)
