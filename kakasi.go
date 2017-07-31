package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type kakasi struct{}

func (k kakasi) generate(text string) (string, error) {

	bodyReq := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyReq)
	defer writer.Close()
	writer.WriteField("state", "output")
	writer.WriteField("text", text)
	writer.WriteField("submit", "Add Furigana !")
	contentType := writer.FormDataContentType()

	resp, err := http.Post("http://furigana.sourceforge.net/cgi-bin/index.cgi", contentType, bodyReq)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	matcher := func(n *html.Node) bool {
		return n.DataAtom == atom.Ruby || n.DataAtom == atom.Rp || n.DataAtom == atom.Rt || (n.Type == html.TextNode && n.Parent.DataAtom == atom.Body)
	}
	children := scrape.FindAll(root, matcher)
	buf := new(bytes.Buffer)
	for _, n := range children {
		html.Render(buf, n)
	}

	return strings.TrimSpace(buf.String()), nil
}
