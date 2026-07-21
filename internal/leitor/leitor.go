package leitor

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/guesant/anp-historico/internal/anp/planilha"
	"github.com/guesant/anp-historico/internal/anp/planilha/tabela"
	"github.com/xuri/excelize/v2"
)

func listarArquivos(pathPlanilhas string) ([]string, error) {
	dir, err := os.ReadDir(pathPlanilhas)

	if err != nil {
		return nil, fmt.Errorf("ler pasta de planilhas: %w", err)
	}

	var arquivos []string

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		arquivos = append(arquivos, path.Join(pathPlanilhas, file.Name()))
	}

	return arquivos, nil
}

func lerPlanilha(arquivo string) error {
	f, err := excelize.OpenFile(arquivo)

	if err != nil {
		return fmt.Errorf("abrir arquivo de planilha: %w", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	propsPasta, err := f.GetWorkbookProps()

	if err != nil {
		return fmt.Errorf("obter atributos da pasta de trabalho: %w", err)
	}

	usarSistema1904 := propsPasta.Date1904 != nil && *propsPasta.Date1904

	sheets := f.GetSheetList()

	for _, sheet := range sheets {
		err2 := lerSheet(f, sheet, usarSistema1904)

		if err2 != nil {
			return err2
		}
	}

	return nil
}

func lerSheet(f *excelize.File, sheet string, usarSistema1904 bool) error {
	rows, err := f.Rows(sheet)

	if err != nil {
		return fmt.Errorf("obter linhas: %w", err)
	}

	defer func(rows *excelize.Rows) {
		_ = rows.Close()
	}(rows)

	detector := planilha.NewDetector()
	var parser *tabela.Parser

	indiceFisico, indiceVirtual, limite := 0, 0, 100

	for rows.Next() {
		indiceFisico++

		colunas, err := rows.Columns(excelize.Options{RawCellValue: true})

		if err != nil {
			return fmt.Errorf("obter colunas da linha: %w", err)
		}

		if len(colunas) == 0 {
			continue
		}

		if parser == nil {
			detector.AnalisarLinha(indiceFisico, colunas)

			if detector.Confirmado() {
				_, err := detector.Formato()

				if err != nil {
					return fmt.Errorf("detectar formato da planilha: %v", err)
				}

				parser, err = tabela.NewParser(colunas, usarSistema1904)

				if err != nil {
					return fmt.Errorf("criar parser para planilha: %v", err)
				}

				continue
			}

			indiceVirtual++

			if indiceVirtual >= limite {
				break
			}

			continue
		}

		registro, err := parser.ProcessarLinha(indiceFisico, colunas)

		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", registro)
	}

	if detector.Confirmado() {
		fmt.Println("boa")
	} else {
		fmt.Println("vish")
	}

	return nil
}

func TodasAsPlanilhas(pathPlanilhas string) error {
	arquivos, err := listarArquivos(pathPlanilhas)

	if err != nil {
		return fmt.Errorf("listar arquivos de planilhas: %v", err)
	}

	considerado := arquivos

	for _, arquivo := range considerado {
		if strings.HasSuffix(arquivo, ".xlsb") || strings.HasSuffix(arquivo, ".xls") {
			log.Printf("WARN: pulando %s pq .xlsb e .xls ainda não com suporte ainda\n", arquivo)
			continue
		}

		err := lerPlanilha(arquivo)

		if err != nil {
			return fmt.Errorf("error de ler planilha:  %v %v", arquivo, err)
		}
	}

	return nil
}
