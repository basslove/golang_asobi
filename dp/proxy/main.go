package main

import "fmt"

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(str string) string
}

type printer struct {
	name string
}

func (p *printer) SetPrinterName(name string) {
	p.name = name
}

func (p *printer) GetPrinterName() string {
	return p.name
}

func (p *printer) Print(str string) string {
	return p.name + ":" + str
}

type PrinterProxy struct {
	Name string
	real *printer
}

func (pp *PrinterProxy) SetPrinterName(name string) {
	if pp.real != nil {
		pp.real.SetPrinterName(name)
	}
	pp.Name = name
}

func (pp *PrinterProxy) GetPrinterName() string {
	return pp.Name
}

func (pp *PrinterProxy) Print(str string) string {
	pp.realize()
	return pp.real.Print(str)
}

func (pp *PrinterProxy) realize() {
	if pp.real == nil {
		pp.real = &printer{pp.Name}
	}
}

func main() {
	var proxy Printable = &PrinterProxy{Name: "A"}
	fmt.Println(proxy.GetPrinterName())

	proxy.SetPrinterName("B")
	fmt.Println(proxy.GetPrinterName())

	fmt.Println(proxy.Print("C"))
}
