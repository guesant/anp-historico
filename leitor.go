package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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

		//panes, err := f.GetPanes(sheet)
		//
		//if err != nil {
		//	return fmt.Errorf("error getting panes: %v", err)
		//}
		//
		//if !panes.Freeze || panes.YSplit == 0 {
		//	return fmt.Errorf("não existem linhas congeladas")
		//}
		//
		//fmt.Printf("arquivo %s freeze %v", arquivo, panes.YSplit)

		//rows, err := f.Rows(sheet)
		//
		//if err != nil {
		//	return fmt.Errorf("error getting rows: %v", err)
		//}
		//
		//curr, limit := 0, 20
		//
		//
		//for rows.Next() {
		//	row, err := rows.Columns()
		//
		//	if err != nil {
		//		return fmt.Errorf("error getting columns: %v", err)
		//	}
		//
		//	if len(row) == 0 {
		//		continue
		//	}
		//
		//	curr++
		//
		//	if curr > limit {
		//		break
		//	}
		//
		//	fmt.Println(row)
		//}
		//
		//if err = rows.Close(); err != nil {
		//	return fmt.Errorf("error closing rows: %v", err)
		//}

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

		fmt.Println()
		fmt.Println()
		fmt.Println()

		err := LerPlanilha(arquivo)

		if err != nil {
			return fmt.Errorf("error de ler planilha:  %v %v", arquivo, err)
		}

	}

	return nil
}
