package main

import "fmt"

type Print interface {
	PrintWeak()
	PrintStrong()
}

var _ Print = &PrintBanner{}

type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	return &Banner{str: str}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%v)", b.str)
}

func (b *Banner) ShowWithAstr() {
	fmt.Printf("**%v**", b.str)
}

type PrintBanner struct {
	Banner *Banner
}

func NewPrintBanner(str string) *PrintBanner {
	return &PrintBanner{Banner: NewBanner(str)}
}

func (pb *PrintBanner) PrintWeak() {
	if pb.Banner != nil {
		pb.Banner.ShowWithParen()
	}
}

func (pb *PrintBanner) PrintStrong() {
	if pb.Banner != nil {
		pb.Banner.ShowWithAstr()
	}
}

func main() {
	fmt.Println("adapter")

	var p Print = NewPrintBanner("hello")
	p.PrintStrong()
	fmt.Println()
	p.PrintWeak()
}
