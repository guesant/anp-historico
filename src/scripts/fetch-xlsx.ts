enum TipoSerie {
  SEMANAL,
  MENSAL
}

enum Abrangencia {
  Brasil,
  Regioes,
  Estados,
  Municipios
}

type Planilha = {
  tipo: TipoSerie;
  abrangencia: Abrangencia;

  url: string;

  de: DateIso;
  ate: DateIso | null;
};

type DateIso = `${number}-${number}-${number}`;

const buildDateIso = (ano: number, mes?: number, dia?: number): DateIso => {
  const mesStr = mes !== undefined ? String(mes).padStart(2, "0") : "01";
  const diaStr = dia !== undefined ? String(dia).padStart(2, "0") : "01";
  return `${ano}-${mesStr}-${diaStr}` as DateIso;
}

/*
2001-2012/semanal-brasil-2004-a-2012.xlsx
2001-2012/semanal-regioes-2004-a-2012.xlsx
2001-2012/semanal-estados-2004-a-2012.xlsx
2001-2012/semanal-municipios-2004-a-2012.xlsb
semanal/semanal-brasil-desde-2013.xlsx
semanal/semanal-regioes-desde-2013.xlsx
semanal/semanal-estados-desde-2013.xlsx
semanal/semanal-municipios-2013-2014.xlsb
semanal/semanal-municipios-2015-a-2017.xlsb
semanal/semanal-municipio-2018-a-2021.xls
semanal/semanal-municipios-2022_a_2023.xlsx
semanal/semanal-municipio-2024-2025.xlsx
semanal/semanal-municipios-2026.xlsx
2001-2012/mensal-brasil-2001-a-2012.xlsx
2001-2012/mensal-regioes-2001-a-2012.xlsx
2001-2012/mensal-estados-2001-a-2012.xlsx
2001-2012/mensal-municipios-2001-a-2012.xlsb
mensal/mensal-brasil-desde-jan2013.xlsx
mensal/mensal-regioes-desde-jan2013.xlsx
mensal/mensal-estados-desde-jan2013.xlsx
mensal/mensal-municipios-2013-a-2015.xlsx
mensal/mensal-municipios-2016-a-2018.xlsx
mensal/mensal-municipios-2019-a-2021.xlsx
mensal/mensal-municipios-jan2022-2025.xlsx
mensal/mensal-municipios-desde-jan2026.xlsx
*/

const base = "https://www.gov.br/anp/pt-br/assuntos/precos-e-defesa-da-concorrencia/precos/precos-revenda-e-de-distribuicao-combustiveis/shlp/";

const namePattern = /^(2001-2012|semanal|mensal)\/(semanal|mensal)-(brasil|regioes|municipios|municipio|estados)-(.+).(xlsx|xlsb|xls)$/;

const normalizarTipo = (tipo: string): TipoSerie => {
  switch (tipo) {
    case "semanal":
      return TipoSerie.SEMANAL;
    case "mensal":
      return TipoSerie.MENSAL;
    default:
      throw new Error(`Tipo de série desconhecido: ${tipo}`);
  }
}

const normalizarAbrangencia = (abrangencia: string): Abrangencia => {
  switch (abrangencia) {
    case "brasil":
      return Abrangencia.Brasil;
    case "regioes":
      return Abrangencia.Regioes;
    case "estados":
      return Abrangencia.Estados;
    case "municipios":
    case "municipio":
      return Abrangencia.Municipios;
    default:
      throw new Error(`Abrangência desconhecida: ${abrangencia}`);
  }
}

const normalizarIntervalo = (intervalo: string): null | { de: DateIso, ate: DateIso | null } => {
  const patternRangeAnos = /^([\d]{4})(-|-a-|_a_)([\d]{4})$/;

  const matchPatternAnos = patternRangeAnos.exec(intervalo);

  if (matchPatternAnos) {
    const [, de, , ate] = matchPatternAnos;

    return {
      de: buildDateIso(Number(de)),
      ate: buildDateIso(Number(ate), 12, 31)
    }
  }

  const patternDesdeAno = /^desde-([\d]{4})$/;

  const matchPatternDesdeAno = patternDesdeAno.exec(intervalo);

  if (matchPatternDesdeAno) {
    const [, de] = matchPatternDesdeAno;

    return {
      de: buildDateIso(Number(de)),
      ate: null
    }
  }

  

  return null;
}

function classificarPlanilha(url: string): Planilha | null {
  const name = url.replace(base, "");
  const match = namePattern.exec(name);



  if (match) {
    const [, periodo, tipo, abrangencia, intervalo] = match;

    const tipoNormalizado = normalizarTipo(tipo);
    const abrangenciaNormalizada = normalizarAbrangencia(abrangencia);
    const intervaloNormalizado = normalizarIntervalo(intervalo);

    if (!intervaloNormalizado) {
      return null;
    }

    return {
      tipo: tipoNormalizado,
      abrangencia: abrangenciaNormalizada,
      url,
      de: intervaloNormalizado.de,
      ate: intervaloNormalizado.ate,
    }
  }

  return null;
}

function* extrairPlanilhas(): Generator<Planilha> {
  const doc = document;

  const links = Array.from(doc.querySelectorAll(`[href*="${base}"]`)).map(a => a.getAttribute("href"));

  for (const link of links) {
    if (!link) continue;
    
    const classificacao = classificarPlanilha(link);

    if (classificacao) {
      yield classificacao;
    }
  }
}
