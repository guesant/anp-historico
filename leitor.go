package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/guesant/anp-historico/internal/anp"
	dd "github.com/guesant/anp-historico/internal/anp/detector"
	"github.com/xuri/excelize/v2"
)

func ListarPlanilhas(pathPlanilhas string) ([]string, error) {
	dir, err := os.ReadDir(pathPlanilhas)

	if err != nil {
		return nil, fmt.Errorf("error reading planilhas: %v", err)
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

type Registro struct {
	DataInicial *time.Time
	DataFinal   *time.Time
	Mes         *time.Time

	Produto   string
	Regiao    string
	Estado    string
	Municipio string

	PostosPesquisados int
	UnidadeMedida     string

	PrecoMedioRevenda  *float64
	DesvioRevenda      *float64
	PrecoMinimoRevenda *float64
	PrecoMaximoRevenda *float64
}

type Parser interface {
	ProcessarLinha(numero int, linha []string) (Registro, error)
}

func NewParser(formato anp.Formato) (*Parser, error) {
	return nil, nil
}

func LerPlanilha(arquivo string) error {
	f, err := excelize.OpenFile(arquivo)

	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetList()

	for _, sheet := range sheets {
		rows, err := f.Rows(sheet)

		if err != nil {
			return fmt.Errorf("error getting rows: %v", err)
		}

		defer rows.Close()

		detector := dd.NewDetector()
		var parser Parser

		indiceFisico, indiceVirtual, limite := 0, 0, 100

		for rows.Next() {
			indiceFisico++

			linha, err := rows.Columns()

			if err != nil {
				return fmt.Errorf("error getting columns: %v", err)
			}

			if len(linha) == 0 {
				continue
			}

			if parser == nil {
				detector.AnalisarLinha(indiceFisico, linha)

				if detector.Confirmado() {
					formato, err := detector.Formato()

					if err != nil {
						return fmt.Errorf("error getting formato: %v", err)
					}

					parser, err := NewParser(formato)

					if err != nil {
						return fmt.Errorf("error creating parser: %v", err)
					}

					break
				}

				indiceVirtual++

				if indiceVirtual >= limite {
					break
				}

				continue
			}

			registro, err := parser.ProcessarLinha(indiceFisico, linha)

			if err != nil {
				return err
			}

			fmt.Printf("%+v\n", registro)
		}

		if detector.Confirmado() {
			fmt.Println("detectou")
		} else {
			fmt.Println("vish")
		}

		if err = rows.Close(); err != nil {
			return fmt.Errorf("error closing rows: %v", err)
		}

	}

	return nil
}

func LerTodasAsPlanilhas(pathPlanilhas string) error {
	arquivos, err := ListarPlanilhas(pathPlanilhas)

	if err != nil {
		return fmt.Errorf("error listar planilhas: %v", err)
	}

	considerado := arquivos

	for _, arquivo := range considerado {
		if strings.HasSuffix(arquivo, ".xlsb") || strings.HasSuffix(arquivo, ".xls") {
			log.Printf("WARN: pulando %s pq .xlsb e .xls ainda não com suporte ainda\n", arquivo)
			continue
		}

		err := LerPlanilha(arquivo)

		if err != nil {
			return fmt.Errorf("error de ler planilha:  %v %v", arquivo, err)
		}
	}

	return nil
}
