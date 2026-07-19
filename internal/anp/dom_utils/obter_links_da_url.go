package dom_utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ObterLinksDaURL(ctx context.Context, paginaURL string) ([]string, error) {
	baseURL, err := url.Parse(paginaURL)

	if err != nil {
		return nil, fmt.Errorf("interpretar URL da página %s", paginaURL)
	}

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

	links, err := ObterLinksDoHtml(response.Body)

	if err != nil {
		return nil, fmt.Errorf("interpretar links da pagina %s", paginaURL)
	}

	return links, nil
}
