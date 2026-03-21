package parser

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type Parser struct {
}

func NewParser(client interface{}) *Parser {
	return &Parser{}
}

func (p *Parser) ExtractLinks(r io.Reader, baseURL string) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			var attrKey string
			switch n.Data {
			case "a", "link", "area":
				attrKey = "href"
			case "img", "script", "iframe", "embed", "source", "track":
				attrKey = "src"
			case "form":
				attrKey = "action"
			}
			if attrKey != "" {
				for _, a := range n.Attr {
					if a.Key == attrKey {
						ref := a.Val
						abs := resolveURL(ref, base)
						if abs != "" {
							links = append(links, abs)
						}
						break
					}
				}
			}

			if n.Data == "img" || n.Data == "source" {
				for _, a := range n.Attr {
					if a.Key == "srcset" {

						parts := strings.Split(a.Val, ",")
						for _, part := range parts {
							fields := strings.Fields(part)
							if len(fields) > 0 {
								abs := resolveURL(fields[0], base)
								if abs != "" {
									links = append(links, abs)
								}
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func resolveURL(ref string, base *url.URL) string {
	u, err := url.Parse(ref)
	if err != nil {
		return ""
	}
	abs := base.ResolveReference(u)

	abs.Fragment = ""
	return abs.String()
}

func IsResource(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	path := u.Path
	ext := strings.ToLower(path[strings.LastIndex(path, ".")+1:])
	switch ext {
	case "css", "js", "png", "jpg", "jpeg", "gif", "svg", "webp", "ico", "ttf", "woff", "woff2", "mp4", "webm":
		return true
	default:
		return false
	}
}
