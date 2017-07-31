package main

import (
	"fmt"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type generator interface {
	generate(text string) (string, error)
}

func main() {
	var opts struct {
		Positional struct {
			Text string `description:"Some japanese text" positional-arg-name:"text" required:"yes"`
		} `positional-args:"yes" required:"yes"`
		Service string `short:"s" long:"service" description:"The furigana generation service. kakasi (http://furigana.sourceforge.net/cgi-bin/index.cgi) or tatoeba (https://tatoeba.org/eng/tools/furigana)" choice:"kakasi" choice:"tatoeba" default:"kakasi"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	var generator generator
	switch opts.Service {
	case "kakasi":
		generator = kakasi{}
	case "tatoeba":
		generator = tatoeba{}
	}

	furigana, err := generator.generate(opts.Positional.Text)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println(furigana)
}
