package dom_utils

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ObterLinksDoNo(no *html.Node) ([]string, error) {
	var links []string

	visitadas := make(map[string]struct{})

	var percorrer func(*html.Node) error

	percorrer = func(no *html.Node) error {
		if no.Type == html.ElementNode && no.Data == "a" {
			href, encontrada := ObterAtributo(no, "href")

			if encontrada {
				href = strings.TrimSpace(href)

				referencia, err := url.Parse(href)

				if err == nil {
					if _, existe := visitadas[referencia.String()]; !existe {
						visitadas[referencia.String()] = struct{}{}
						links = append(links, href)
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
