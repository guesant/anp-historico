package anp

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const PaginaSerieHistoricaURL = "https://www.gov.br/anp/pt-br/assuntos/precos-e-defesa-da-concorrencia/precos/precos-revenda-e-de-distribuicao-combustiveis/serie-historica-do-levantamento-de-precos"

func obterAtributo(no *html.Node, nome string) (string, bool) {
	for _, atributo := range no.Attr {
		if atributo.Key == nome {
			return atributo.Val, true
		}
	}

	return "", false
}

func ExtrairPlanilhas(
	ctx context.Context,
	paginaURL string,
) ([]Planilha, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		paginaURL,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("criar requisição: %w", err)
	}

	req.Header.Set(
		"User-Agent",
		"anp-historico/1.0",
	)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("baixar página: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"servidor retornou status %s",
			response.Status,
		)
	}

	documento, err := html.Parse(response.Body)

	if err != nil {
		return nil, fmt.Errorf("interpretar HTML: %w", err)
	}

	baseURL, err := url.Parse(paginaURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar URL da página %s", paginaURL)
	}

	planilhas := make([]Planilha, 0)
	visitadas := make(map[string]struct{})

	var percorrer func(*html.Node) error

	percorrer = func(no *html.Node) error {
		if no.Type == html.ElementNode && no.Data == "a" {
			href, encontrada := obterAtributo(no, "href")

			if encontrada {
				href = strings.TrimSpace(href)

				referencia, err := url.Parse(href)

				if err == nil {
					urlAbsoluta := baseURL.ResolveReference(referencia).String()

					planilha, reconhecida, err := ClassificarPlanilha(urlAbsoluta)

					if err != nil {
						return fmt.Errorf("classificar %q: %w", urlAbsoluta, err)
					}

					if reconhecida {
						if _, existe := visitadas[urlAbsoluta]; !existe {
							visitadas[urlAbsoluta] = struct{}{}

							planilhas = append(planilhas, planilha)
						}
					}
				}
			}
		}

		for filho := no.FirstChild; filho != nil; filho = filho.NextSibling {
			if err := percorrer(filho); err != nil {
				return err
			}
		}

		return nil
	}

	if err := percorrer(documento); err != nil {
		return nil, err
	}

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Tipo) < string(planilhas[j].Tipo)
	})

	sort.Slice(planilhas, func(i, j int) bool {
		return string(planilhas[i].Abrangencia) < string(planilhas[j].Abrangencia)
	})

	sort.Slice(planilhas, func(i, j int) bool {
		return planilhas[i].De < planilhas[j].De
	})

	return planilhas, nil
}
