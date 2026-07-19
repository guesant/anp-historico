package dom_utils

import "golang.org/x/net/html"

func ObterAtributo(no *html.Node, nome string) (string, bool) {
	for _, atributo := range no.Attr {
		if atributo.Key == nome {
			return atributo.Val, true
		}
	}

	return "", false
}
