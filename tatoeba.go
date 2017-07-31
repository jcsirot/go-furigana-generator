package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type tatoeba struct{}

func (t tatoeba) generate(text string) (string, error) {

	resp, err := http.Get(fmt.Sprintf("https://tatoeba.org/eng/tools/furigana?query=%s", text))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	matcher := func(n *html.Node) bool {
		return scrape.Attr(n, "id") == "conversion"
	}

	div, _ := scrape.Find(root, matcher)
	matcher = func(n *html.Node) bool {
		return n.DataAtom != atom.Span && n.Parent == div
	}
	children := scrape.FindAll(div, matcher)
	buf := new(bytes.Buffer)
	for _, n := range children {
		html.Render(buf, n)
	}
	return buf.String(), nil
}
