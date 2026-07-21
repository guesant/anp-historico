package links_planilhas

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func obterLinksDaURL(ctx context.Context, paginaURL string) ([]string, error) {
	baseURL, err := url.Parse(paginaURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar URL da página %s", paginaURL)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		baseURL.String(),
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

	links, err := obterLinksDoHtml(response.Body, baseURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", baseURL.String())
	}

	return links, nil
}

func obterLinksDoHtml(htmlContent io.Reader, baseURL *url.URL) ([]string, error) {
	documento, err := html.Parse(htmlContent)

	if err != nil {
		return nil, fmt.Errorf("interpretar HTML: %w", err)
	}

	links, err := obterLinksDoNo(documento, baseURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", err)
	}

	return links, nil
}

func obterLinksDoNo(no *html.Node, baseURL *url.URL) ([]string, error) {
	var links []string

	visitadas := make(map[string]struct{})

	var percorrer func(*html.Node) error

	percorrer = func(no *html.Node) error {
		if no.Type == html.ElementNode && no.Data == "a" {
			href, encontrada := obterAtributoDoNo(no, "href")

			if encontrada {
				href = strings.TrimSpace(href)

				referencia, err := url.Parse(href)

				if err == nil {
					var urlAbsoluta string

					if baseURL != nil {
						urlAbsoluta = baseURL.ResolveReference(referencia).String()
					} else {
						urlAbsoluta = referencia.String()
					}

					_, existe := visitadas[urlAbsoluta]

					if !existe {
						visitadas[urlAbsoluta] = struct{}{}
						links = append(links, urlAbsoluta)
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

	if err := percorrer(no); err != nil {
		return nil, err
	}

	return links, nil
}

func obterAtributoDoNo(no *html.Node, nome string) (string, bool) {
	for _, atributo := range no.Attr {
		if atributo.Key == nome {
			return atributo.Val, true
		}
	}

	return "", false
}
