package dom_utils

import (
	"fmt"
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func ObterLinksDoHtml(htmlContent io.Reader, baseURL *url.URL) ([]string, error) {
	documento, err := html.Parse(htmlContent)

	if err != nil {
		return nil, fmt.Errorf("interpretar HTML: %w", err)
	}

	links, err := ObterLinksDoNo(documento, baseURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", err)
	}

	return links, nil
}
