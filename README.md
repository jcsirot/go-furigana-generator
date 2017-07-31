# go-furigana-generator
HTML furagana generator based on web services.

The following furigana generation service are available :

- [Furigana Generator with HTML](http://furigana.sourceforge.net/cgi-bin/index.cgi) using [Kakasi](http://kakasi.namazu.org/) library
- [Tatoeba](https://tatoeba.org/eng/tools/furigana)

## Install

    go get -u github.com/jcsirot/go-furigana-generator

## Build

    go build -o go-furigana-generator

## Usage

``` plain
Usage:
  go-furigana-generator [OPTIONS] text

Application Options:
  -s, --service=[kakasi|tatoeba] The furigana generation service. kakasi
                                 (http://furigana.sourceforge.net/cgi-bin/index.cgi) or
                                 tatoeba (https://tatoeba.org/eng/tools/furigana) (default:
                                 kakasi)

Help Options:
  -h, --help                     Show this help message

Arguments:
  text:                          Some japanese text
```

## Usage Example

    go-furigana-generator -s kakasi -- 学校は家から遠いの？

    <ruby><rb>学校</rb><rp>(</rp><rt>がっこう</rt><rp>)</rp></ruby>は<ruby><rb>家</rb><rp>(</rp><rt>いえ</rt><rp>)</rp></ruby>から<ruby><rb>遠</rb><rp>(</rp><rt>とお</rt><rp>)</rp></ruby>いの？
