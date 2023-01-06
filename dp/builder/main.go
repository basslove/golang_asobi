package main

import (
	"fmt"
	"os"
	"strings"
)

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

var _ Builder = &TextBuilder{}
var _ Builder = &HtmlBuilder{}

type TextBuilder struct {
	StringBuilder []string
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{StringBuilder: make([]string, 0)}
}

func (tb *TextBuilder) MakeTitle(title string) {
	tb.StringBuilder = append(tb.StringBuilder, "================================")
	tb.StringBuilder = append(tb.StringBuilder, "『")
	tb.StringBuilder = append(tb.StringBuilder, title)
	tb.StringBuilder = append(tb.StringBuilder, "』")
}

func (tb *TextBuilder) MakeString(str string) {
	tb.StringBuilder = append(tb.StringBuilder, "■")
	tb.StringBuilder = append(tb.StringBuilder, str)
	tb.StringBuilder = append(tb.StringBuilder, "\n\n")
}

func (tb *TextBuilder) MakeItems(items []string) {
	for _, s := range items {
		tb.StringBuilder = append(tb.StringBuilder, "・")
		tb.StringBuilder = append(tb.StringBuilder, s)
		tb.StringBuilder = append(tb.StringBuilder, "\n")
	}
	tb.StringBuilder = append(tb.StringBuilder, "\n")
}

func (tb *TextBuilder) Close() {
	tb.StringBuilder = append(tb.StringBuilder, "===============================")
	tb.StringBuilder = append(tb.StringBuilder, "===============================")
}

type HtmlBuilder struct {
	filename      string
	StringBuilder []string
}

func NewHtmlBuilder() *HtmlBuilder {
	return &HtmlBuilder{filename: "untitle.html", StringBuilder: make([]string, 0)}
}

func (hb *HtmlBuilder) MakeTitle(title string) {
	hb.StringBuilder = append(hb.StringBuilder, "<!DOCTYPE html>\n")
	hb.StringBuilder = append(hb.StringBuilder, "<html>\n")
	hb.StringBuilder = append(hb.StringBuilder, "<head><title>\n")
	hb.StringBuilder = append(hb.StringBuilder, title)
	hb.StringBuilder = append(hb.StringBuilder, "</head></title>\n")
	hb.StringBuilder = append(hb.StringBuilder, "<body>\n")
	hb.StringBuilder = append(hb.StringBuilder, "<h1>\n")
	hb.StringBuilder = append(hb.StringBuilder, title)
	hb.StringBuilder = append(hb.StringBuilder, "</h1>\n\n")

}

func (hb *HtmlBuilder) MakeString(str string) {
	hb.StringBuilder = append(hb.StringBuilder, "<p>")
	hb.StringBuilder = append(hb.StringBuilder, str)
	hb.StringBuilder = append(hb.StringBuilder, "</p>\n\n")
}

func (hb *HtmlBuilder) MakeItems(items []string) {
	hb.StringBuilder = append(hb.StringBuilder, "<ul>")
	for _, s := range items {
		hb.StringBuilder = append(hb.StringBuilder, "<li>")
		hb.StringBuilder = append(hb.StringBuilder, s)
		hb.StringBuilder = append(hb.StringBuilder, "</li>\n")
	}
	hb.StringBuilder = append(hb.StringBuilder, "</ul>\n")
}

func (hb *HtmlBuilder) Close() {
	hb.StringBuilder = append(hb.StringBuilder, "</body>")
	hb.StringBuilder = append(hb.StringBuilder, "</html>\n")

	data := []byte(strings.Join(hb.StringBuilder, ","))
	err := os.WriteFile(hb.filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("greeting")
	d.builder.MakeString("一般的交渉")
	d.builder.MakeItems([]string{"hello", "hi."})
	d.builder.MakeString("時間的交渉")
	d.builder.MakeItems([]string{"good morning", "good afternoon", "good evening"})
	d.builder.Close()
}

func main() {
	fmt.Println("builder")

	moji := "html"
	if moji == "text" {
		tb := NewTextBuilder()
		d := NewDirector(tb)
		d.Construct()
		fmt.Println(strings.Join(tb.StringBuilder, ","))
	}
	if moji == "html" {
		hb := NewHtmlBuilder()
		d := NewDirector(hb)
		d.Construct()
		fmt.Println(hb.filename)
	}
}
