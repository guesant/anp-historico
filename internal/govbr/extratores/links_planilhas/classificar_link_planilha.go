package links_planilhas

func ClassificarLinkPlanilha(rawURL string) (LinkPlanilha, bool, error) {
	nome, reconhecida := nomeRelativo(rawURL)

	if !reconhecida {
		return LinkPlanilha{}, false, nil
	}

	match := nomePlanilhaPattern.FindStringSubmatch(nome)
	if match == nil {
		return LinkPlanilha{}, false, nil
	}

	tipo, err := normalizarTipo(match[1])
	if err != nil {
		return LinkPlanilha{}, true, nil
	}

	abrangencia, err := normalizarAbrangencia(match[2])
	if err != nil {
		return LinkPlanilha{}, true, err
	}

	referenciaIntervalo, err := normalizarIntervalo(match[3])
	if err != nil {
		return LinkPlanilha{}, true, err
	}

	return LinkPlanilha{
		Tipo:                tipo,
		Abrangencia:         abrangencia,
		URL:                 rawURL,
		ReferenciaIntervalo: referenciaIntervalo,
	}, true, nil
}
