package dom_utils

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func ObterLinksDoHtml(htmlContent io.Reader) ([]string, error) {
	documento, err := html.Parse(htmlContent)

	if err != nil {
		return nil, fmt.Errorf("interpretar HTML: %w", err)
	}

	links, err := ObterLinksDoNo(documento)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", htmlContent)
	}

	return links, nil
}
