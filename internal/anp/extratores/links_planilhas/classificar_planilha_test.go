package links_planilhas

import (
	"testing"

	"github.com/guesant/anp-historico/internal/anp"
)

func TestClassificarPlanilhasConhecidas(t *testing.T) {
	testes := []struct {
		nome        string
		tipo        anp.TipoSerie
		abrangencia anp.Abrangencia
		de          string
		ate         string
	}{
		{
			nome:        "2001-2012/semanal-brasil-2004-a-2012.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaBrasil,
			de:          "2004-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/semanal-regioes-2004-a-2012.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaRegioes,
			de:          "2004-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/semanal-estados-2004-a-2012.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaEstados,
			de:          "2004-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/semanal-municipios-2004-a-2012.xlsb",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2004-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "semanal/semanal-brasil-desde-2013.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaBrasil,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "semanal/semanal-regioes-desde-2013.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaRegioes,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "semanal/semanal-estados-desde-2013.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaEstados,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "semanal/semanal-municipios-2013-2014.xlsb",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2013-01-01",
			ate:         "2014-12-31",
		},
		{
			nome:        "semanal/semanal-municipios-2015-a-2017.xlsb",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2015-01-01",
			ate:         "2017-12-31",
		},
		{
			nome:        "semanal/semanal-municipio-2018-a-2021.xls",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2018-01-01",
			ate:         "2021-12-31",
		},
		{
			nome:        "semanal/semanal-municipios-2022_a_2023.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2022-01-01",
			ate:         "2023-12-31",
		},
		{
			nome:        "semanal/semanal-municipio-2024-2025.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2024-01-01",
			ate:         "2025-12-31",
		},
		{
			nome:        "semanal/semanal-municipios-2026.xlsx",
			tipo:        anp.TipoSerieSemanal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2026-01-01",
			ate:         "2026-12-31",
		},
		{
			nome:        "2001-2012/mensal-brasil-2001-a-2012.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaBrasil,
			de:          "2001-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/mensal-regioes-2001-a-2012.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaRegioes,
			de:          "2001-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/mensal-estados-2001-a-2012.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaEstados,
			de:          "2001-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "2001-2012/mensal-municipios-2001-a-2012.xlsb",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2001-01-01",
			ate:         "2012-12-31",
		},
		{
			nome:        "mensal/mensal-brasil-desde-jan2013.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaBrasil,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "mensal/mensal-regioes-desde-jan2013.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaRegioes,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "mensal/mensal-estados-desde-jan2013.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaEstados,
			de:          "2013-01-01",
			ate:         "",
		},
		{
			nome:        "mensal/mensal-municipios-2013-a-2015.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2013-01-01",
			ate:         "2015-12-31",
		},
		{
			nome:        "mensal/mensal-municipios-2016-a-2018.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2016-01-01",
			ate:         "2018-12-31",
		},
		{
			nome:        "mensal/mensal-municipios-2019-a-2021.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2019-01-01",
			ate:         "2021-12-31",
		},
		{
			nome:        "mensal/mensal-municipios-jan2022-2025.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2022-01-01",
			ate:         "2025-12-31",
		},
		{
			nome:        "mensal/mensal-municipios-desde-jan2026.xlsx",
			tipo:        anp.TipoSerieMensal,
			abrangencia: anp.AbrangenciaMunicipios,
			de:          "2026-01-01",
			ate:         "",
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			url := anp.BasePlanilhasURL + teste.nome

			resultado, reconhecida, err := ClassificarPlanilha(url)

			if err != nil {
				t.Fatalf("ClassificarPlanilha() retornou erro: %v", err)
			}

			if !reconhecida {
				t.Fatal("ClassificarPlanilha() não reconheceu a URL")
			}

			if resultado.Tipo != teste.tipo {
				t.Errorf(
					"Abrangência = %q; esperado = %q",
					resultado.Abrangencia,
					teste.abrangencia,
				)
			}

			if string(resultado.AnoInicioAproximado) != teste.de {
				t.Errorf(
					"De = %q; esperado %q",
					resultado.AnoInicioAproximado,
					teste.de,
				)
			}

			if teste.ate == "" {
				if resultado.AnoTerminoAproximado != nil {
					t.Errorf("Até = %q; esperado nil", *resultado.AnoTerminoAproximado)
				}

				return
			}

			if resultado.AnoTerminoAproximado == nil {
				t.Fatalf(
					"Ate = nil; esperado %q",
					teste.ate,
				)
			}

			if string(*resultado.AnoTerminoAproximado) != teste.ate {
				t.Errorf(
					"Ate = %q; esperado %q",
					*resultado.AnoTerminoAproximado,
					teste.ate,
				)
			}

		})
	}
}

func TestClassificarPlanilhaIgnoraURLDesconhecida(t *testing.T) {
	_, reconhecida, err := ClassificarPlanilha(
		"https://example.com/arquivo.xlsx",
	)

	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if reconhecida {
		t.Fatal("URL desconhecida foi reconhecida")
	}
}

func TestClassificarPlanilhaDetectaIntervaloDesconhecido(t *testing.T) {
	_, reconhecida, err := ClassificarPlanilha(
		anp.BasePlanilhasURL +
			"semanal/semanal-brasil-periodo-desconhecido.xlsx",
	)

	if !reconhecida {
		t.Fatal("o formato geral da URL deveria ter sido reconhecido")
	}

	if err == nil {
		t.Fatal("era esperado um erro de intervalo desconhecido")
	}
}
